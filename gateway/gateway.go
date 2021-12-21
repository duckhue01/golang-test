package gateway

import (
  "context"
  "flag"
  "net/http"
  "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
  "google.golang.org/grpc"
  gw "github.com/duckhue01/golang_test/proto/v1"  
  // "github.com/golang/glog"
)

var (
  // command-line options:
  // gRPC server endpoint
  grpcServerEndpoint = flag.String("grpc-server-endpoint",  "localhost:4040", "gRPC server endpoint")
)

func RunGateWay() error {
  ctx := context.Background()
  ctx, cancel := context.WithCancel(ctx)
  defer cancel()

  // Register gRPC server endpoint
  // Note: Make sure the gRPC server is running properly and accessible
  mux := runtime.NewServeMux()
  opts := []grpc.DialOption{grpc.WithInsecure()}
  err := gw.RegisterTodosServiceHandlerFromEndpoint(ctx, mux,  *grpcServerEndpoint, opts)
  if err != nil {
    return err
  }

  // Start HTTP server (and proxy calls to gRPC server endpoint)
  return http.ListenAndServe(":8081", mux)
}

// func main() {
//   flag.Parse()
//   defer glog.Flush()

//   if err := run(); err != nil {
//     glog.Fatal(err)
//   }
// }