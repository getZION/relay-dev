package lightning

import (
	"context"
	"io/ioutil"
	"os"

	rp "github.com/getzion/relampago"
	"github.com/kelseyhightower/envconfig"
	"github.com/lightningnetwork/lnd/lnrpc"
	"github.com/lightningnetwork/lnd/lnrpc/routerrpc"
	"github.com/lightningnetwork/lnd/macaroons"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	macaroon "gopkg.in/macaroon.v2"
)

type ConnectionParams struct {
	Host         string `envconfig:"LND_HOST"`
	CertPath     string `envconfig:"LND_CERT_PATH"`
	MacaroonPath string `envconfig:"LND_MACAROON_PATH"`
}

type LndWallet struct {
	Params ConnectionParams

	Conn      *grpc.ClientConn
	Lightning lnrpc.LightningClient
	Router    routerrpc.RouterClient

	invoiceStatusListeners []chan rp.InvoiceStatus
	paymentStatusListeners []chan rp.PaymentStatus
}

var log = zerolog.New(os.Stderr).Output(zerolog.ConsoleWriter{Out: os.Stdout})
var wallet LndWallet

func Connect() (*LndWallet, error) {
	var params ConnectionParams
	envconfig.Process("", &params)
	log.Info().Str("LND_HOST", params.Host).Msg("Connecting to LND...")

	var dialOpts []grpc.DialOption

	// TLS
	tls, err := credentials.NewClientTLSFromFile(params.CertPath, "")
	if err != nil {
		return nil, err
	}
	dialOpts = append(dialOpts, grpc.WithTransportCredentials(tls))

	// Macaroon Auth
	macBytes, err := ioutil.ReadFile(params.MacaroonPath)
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

	log.Info().Uint64("Balance", balance).Str("Pubkey", pubkey).Msg("Connected to LND.")

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
