package main

import (
	"fmt"
	"net"
	"net/http"
	"time"

	. "github.com/getzion/relay/api"
	. "github.com/getzion/relay/gen/proto/zion/v1"
	"github.com/getzion/relay/lightning"
	"github.com/getzion/relay/storage"
	. "github.com/getzion/relay/utils"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	gorm "github.com/jinzhu/gorm"
	"google.golang.org/grpc"
)

var db *gorm.DB
var err error

type grpcMultiplexer struct {
	*grpcweb.WrappedGrpcServer
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func main() {
	// Connect to LND
	lightning.Connect()

	// Connect to MySQL database
	db, err = storage.ConnectDB()
	if err != nil {
		Log.Fatal().Err(err)
		// panic(err)
	}

	// Start listening on a TCP Port
	lis, err := net.Listen("tcp", "127.0.0.1:9990")
	if err != nil {
		Log.Fatal().Err(err)
	}

	// We need to tell the code WHAT TO do on each request, ie. The business logic.
	// In GRPC cases, the Server is acutally just an Interface
	// So we need a struct which fulfills the server interface
	// see server.go
	apiserver := grpc.NewServer()
	RegisterNodeInfoServiceServer(apiserver, &NodeinfoService{})
	// Start serving in a goroutine to not block
	go func() {
		Log.Fatal().Err(apiserver.Serve(lis))
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
	// Serve the webapp over TLS
	Log.Fatal().Err(srv.ListenAndServe())
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
