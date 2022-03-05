package main

import (
	"fmt"
	"io/fs"
	"net"
	"net/http"
	"time"

	. "github.com/getzion/relay/api"
	. "github.com/getzion/relay/gen/proto/zion/v1"
	"github.com/getzion/relay/lightning"
	. "github.com/getzion/relay/utils"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

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

func init_gRPC() {
	grpcServer := grpc.NewServer()

	RegisterNodeInfoServiceServer(grpcServer, &NodeinfoService{})

	wrappedServer := grpcweb.WrapServer(grpcServer,
		grpcweb.WithAllowNonRootResource(true),
	)
	handler := func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Set("Access-Control-Allow-Origin", "*")
		resp.Header().Set("Access-Control-Allow-Headers", "x-grpc-web,content-type")
		if wrappedServer.IsGrpcWebRequest(req) {
			Log.Info().Msg("Yes is gRPC web request")
			wrappedServer.ServeHTTP(resp, req)
		} else {
			Log.Info().Str("donno", req.RequestURI).Msg("No is not gRPC web request")
			if req.RequestURI == "/" {
				// fmt.Fprint(resp, "hello WORLD")
				distFS, _ := fs.Sub(nextFS, "ui/dist")
				http.Handle("/", http.FileServer(http.FS(distFS)))
				// resp.WriteHeader(200)

			}
		}
	}

	port := 5000

	httpServer := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: http.HandlerFunc(handler),
	}

	Log.Info().Int("port", port).Msg("gRPC listening")
	if err = httpServer.ListenAndServe(); err != nil {
		grpclog.Fatalf("failed starting http server: %v", err)
	}
}
