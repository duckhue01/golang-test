package main

import (
	"context"
	"fmt"
	"time"

	proto "github.com/duckhue01/golang_test/proto/v1"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	t := time.Now().In(time.UTC)
	createAt := timestamppb.New(t)

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

	res1, err := client.GetOne(context.Background(), &proto.GetOneRequest{
		Api: "v1",
		Id:  2,
	})

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res1)

	// res2, err := client.GetAll(context.Background(), &proto.GetAllRequest{
	// 	Api: "v1",
	// })

	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(res2)

	res3, err := client.Update(context.Background(), &proto.UpdateRequest{
		Api: "v1",
		Todo: &proto.Todo{
			Id:          1,
			Title:       "duckhue01",
			Description: "duckhue01",
			CreateAt:    createAt,
			UpdateAt:    createAt,
			IsDone:      true,
		},
	})

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res3)

	res4, err := client.Delete(context.Background(), &proto.DeleteRequest{
		Api: "v1",
		Id:  1,
	})

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res4)
}
