package main

import (
	"net"
	"net/http"
	"time"

	"github.com/getzion/relay/api/identityhub"
	"github.com/getzion/relay/api/nodeinfo"
	"github.com/getzion/relay/api/schema"
	"github.com/getzion/relay/api/storage"
	"github.com/getzion/relay/api/validator"
	hub "github.com/getzion/relay/gen/proto/identityhub/v1"
	zion "github.com/getzion/relay/gen/proto/zion/v1"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcMultiplexer struct {
	*grpcweb.WrappedGrpcServer
}

func main() {

	validator.InitValidator()

	connection, err := storage.NewStorage("mysql")
	if err != nil {
		logrus.Panic(err)
	}

	schemaManager := schema.NewSchemaManager(connection)

	// Initialize gRPC server
	init_gRPC(schemaManager)
}

func init_gRPC(schemaManager *schema.SchemaManager) {
	// Start listening on a TCP Port
	lis, err := net.Listen("tcp", "127.0.0.1:9990")
	if err != nil {
		logrus.Panic(err)
	}

	// We need to tell the code WHAT TO do on each request, ie. The business logic.
	// In GRPC cases, the Server is actually just an Interface
	// So we need a struct which fulfills the server interface
	// see server.go
	apiserver := grpc.NewServer()
	identityHub := identityhub.InitIdentityHubService(schemaManager)
	nodeinfo := nodeinfo.InitNewNodeInfoService()

	zion.RegisterNodeInfoServiceServer(apiserver, nodeinfo)
	hub.RegisterHubRequestServiceServer(apiserver, identityHub)
	reflection.Register(apiserver)
	// Start serving in a goroutine to not block
	go func() {
		logrus.Fatal(apiserver.Serve(lis))
	}()

	// Wrap the GRPC Server in grpc-web and also host the UI
	grpcWebServer := grpcweb.WrapServer(apiserver)
	// Lets put the wrapped grpc server in our multiplexer struct so
	// it can reach the grpc server in its handler
	multiplex := grpcMultiplexer{
		grpcWebServer,
	}
	// We need a http router
	r := http.NewServeMux()
	// Load the static webpage with a http fileserver
	webapp := http.FileServer(http.Dir("ui/dist"))
	// Host the Web Application at /, and wrap it in the GRPC Multiplexer
	// This allows grpc requests to transfer over HTTP1. then be
	// routed by the multiplexer
	r.Handle("/", multiplex.Handler(webapp))
	// Create a HTTP server and bind the router to it, and set wanted address
	srv := &http.Server{
		Handler:      r,
		Addr:         "localhost:5000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logrus.Infof("identityhub server started: %s", lis.Addr().String())
	// Serve the webapp over TLS
	logrus.Error(srv.ListenAndServe())
}

// Handler is used to route requests to either grpc or to regular http
func (m *grpcMultiplexer) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if m.IsGrpcWebRequest(r) {
			m.ServeHTTP(w, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}
