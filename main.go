package main

import (
	"fmt"
	"net/http"

	"github.com/getzion/relay/api"
	. "github.com/getzion/relay/api"
	hub "github.com/getzion/relay/gen/proto/identityhub/v1"
	. "github.com/getzion/relay/gen/proto/zion/v1"
	"github.com/getzion/relay/lightning"
	"github.com/getzion/relay/storage"
	. "github.com/getzion/relay/utils"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	gorm "github.com/jinzhu/gorm"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

var db *gorm.DB
var err error

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func main() {
	Log.Info().Msg("Initializing relay")

	// Connect to LND
	lightning.Connect()

	// Connect to MySQL database
	db, err = storage.Connect("root@tcp(localhost:3306)/relay3")
	if err != nil {
		panic(err)
	}

	// Initialize gRPC server
	init_gRPC()
}

func init_gRPC() {
	grpcServer := grpc.NewServer()

	identityHub := api.InitIdentityHubService()

	RegisterNodeInfoServiceServer(grpcServer, &NodeinfoService{})
	RegisterCommunitiesServiceServer(grpcServer, &CommunitiesServiceDefaultServer{DB: db})
	hub.RegisterHubRequestServiceServer(grpcServer, identityHub)

	wrappedServer := grpcweb.WrapServer(grpcServer)
	handler := func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Set("Access-Control-Allow-Origin", "*")
		resp.Header().Set("Access-Control-Allow-Headers", "x-grpc-web,content-type")
		wrappedServer.ServeHTTP(resp, req)
	}

	httpServer := http.Server{
		Addr:    fmt.Sprintf(":%d", 9090),
		Handler: http.HandlerFunc(handler),
	}

	go func() {
		Log.Info().Msg("HTTP - Listening on port 5000")
		http.HandleFunc("/", helloWorld)
		http.ListenAndServe(":5000", nil)
	}()

	Log.Info().Msg("gRPC - Listening on port 9090")
	if err = httpServer.ListenAndServe(); err != nil {
		grpclog.Fatalf("failed starting http server: %v", err)
	}
}
