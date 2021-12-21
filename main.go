package main

import (
	"net"

	// "github.com/duckhue01/golang_test/gateway"
	"github.com/duckhue01/golang_test/db"
	proto "github.com/duckhue01/golang_test/proto/v1"
	"github.com/duckhue01/golang_test/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	proto.RegisterTodosServiceServer(srv, services.NewTodosService(db.NewDatabaseStore()))
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}
	// gateway.RunGateWay()
}
