package lightning

import (
	"context"
	"io/ioutil"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	rp "github.com/getzion/relampago"
	. "github.com/getzion/relay/utils"
	"github.com/kelseyhightower/envconfig"
	"github.com/lightningnetwork/lnd/lnrpc"
	"github.com/lightningnetwork/lnd/lnrpc/routerrpc"
	"github.com/lightningnetwork/lnd/macaroons"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	macaroon "gopkg.in/macaroon.v2"
)

type ConnectionParams struct {
	Host            string `envconfig:"LND_HOST"`
	KeyBucketRegion string `envconfig:"S3_KEY_BUCKET_REGION"`
	KeyBucketName   string `envconfig:"S3_KEY_BUCKET_NAME"`
	MacaroonKey     string `envconfig:"S3_MACAROON_KEY"`
	CertKey         string `envconfig:"S3_CERT_KEY"`
}

type LndWallet struct {
	Params ConnectionParams

	Conn      *grpc.ClientConn
	Lightning lnrpc.LightningClient
	Router    routerrpc.RouterClient

	invoiceStatusListeners []chan rp.InvoiceStatus
	paymentStatusListeners []chan rp.PaymentStatus
}

var wallet LndWallet
var params ConnectionParams

func Connect() (*LndWallet, error) {
	envconfig.Process("", &params)
	LoadLNDCredentials()

	Log.Info().Str("LND_HOST", params.Host).Msg("Connecting to LND...")

	var dialOpts []grpc.DialOption

	// TLS
	tls, err := credentials.NewClientTLSFromFile("tls.cert", "")
	if err != nil {
		log.Fatal().Err(err)
		return nil, err
	}
	dialOpts = append(dialOpts, grpc.WithTransportCredentials(tls))

	// Macaroon Auth
	macBytes, err := ioutil.ReadFile("admin.macaroon")
	if err != nil {
		return nil, err
	}
	m := &macaroon.Macaroon{}
	err = m.UnmarshalBinary(macBytes)
	if err != nil {
		return nil, err
	}
	creds, err := macaroons.NewMacaroonCredential(m)
	if err != nil {
		return nil, err
	}
	dialOpts = append(dialOpts, grpc.WithPerRPCCredentials(creds))
	dialOpts = append(dialOpts, grpc.WithBlock())

	// Connect
	conn, err := grpc.Dial(params.Host, dialOpts...)
	if err != nil {
		return nil, err
	}
	ln := lnrpc.NewLightningClient(conn)
	router := routerrpc.NewRouterClient(conn)

	l := &LndWallet{
		Params:    params,
		Conn:      conn,
		Lightning: ln,
		Router:    router,
	}
	wallet = *l

	balance, pubkey := GetNodeInfo()

	Log.Info().Uint64("Balance", balance).Str("Pubkey", pubkey).Msg("Connected to LND.")

	return l, nil
}

func GetNodeInfo() (balance uint64, pubkey string) {
	res, err := wallet.Lightning.ChannelBalance(context.Background(), &lnrpc.ChannelBalanceRequest{})
	if err != nil {
		log.Fatal().Err(err).Msg("Error calling ChannelBalance")
		panic(err)
	}
	balance = uint64(res.LocalBalance.Sat)

	info, err2 := wallet.Lightning.GetInfo(context.Background(), &lnrpc.GetInfoRequest{})
	if err2 != nil {
		log.Fatal().Err(err2).Msg("Error calling GetInfo")
		panic(err2)
	}
	pubkey = info.IdentityPubkey

	return balance, pubkey
}

func LoadLNDCredentials() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(params.KeyBucketRegion)},
	)
	if err != nil {
		Log.Fatal().Err(err)
	}

	macaroonFile, err := os.Create("admin.macaroon")
	if err != nil {
		Log.Fatal().Err(err)
	}
	downloader := s3manager.NewDownloader(sess)
	downloaded, err := downloader.Download(macaroonFile, &s3.GetObjectInput{
		Bucket: aws.String(params.KeyBucketName),
		Key:    aws.String(params.MacaroonKey),
	})
	if err != nil {
		Log.Fatal().Err(err)
	}

	Log.Info().Int64("bytes", downloaded).Msg("Loaded macaroon file.")

	file, err := os.Create("tls.cert")
	if err != nil {
		Log.Fatal().Err(err)
	}
	defer file.Close()
	downloaded2, err := downloader.Download(file, &s3.GetObjectInput{
		Bucket: aws.String(params.KeyBucketName),
		Key:    aws.String(params.CertKey),
	})

	if err != nil {
		Log.Fatal().Err(err)
	}

	Log.Info().Int64("bytes", downloaded2).Msg("Loaded cert file.")
}

type LndCredentials struct {
	macaroon     []byte
	macaroonFile *os.File
	cert         []byte
	certFile     *os.File
}
