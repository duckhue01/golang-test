package main

import (
	"context"
	"fmt"
	"time"

	proto "github.com/duckhue01/golang_test/proto/v1"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	t := time.Now().In(time.UTC)
	createAt, _ := ptypes.TimestampProto(t)

	client := proto.NewTodosServiceClient(conn)
	res, err := client.Add(context.Background(), &proto.AddRequest{
		Api: "v1",
		Todo: &proto.Todo{
			Id:          1,
			Title:       "asdasd",
			Description: "asdasd",
			CreateAt:    createAt,
			UpdateAt:    createAt,
			IsDone:      false,
		},
	})
	

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}
