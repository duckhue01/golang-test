package main

import (
	"context"
	"flag"
	"net/http"

	gw "github.com/duckhue01/golang_test/proto/v2"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)


// make sure to change this host address to grpc when deploy to docker
var (
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "grpc:4040", "gRPC server endpoint")
)

func runGateWay() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register grpc server endpoint
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gw.RegisterTodosServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(":8080", mux)
}

func main() {
	flag.Parse()
	defer glog.Flush()
	if err := runGateWay(); err != nil {
		glog.Fatal(err)
	}
}
