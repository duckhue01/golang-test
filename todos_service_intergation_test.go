package main

import (
	"context"
	"fmt"
	"testing"
	"time"

	proto "github.com/duckhue01/golang_test/proto/v2"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const apiVer = "v2"

// this test need to run todos service first
var conn, _ = grpc.Dial("localhost:4040", grpc.WithInsecure())
var client = proto.NewTodosServiceClient(conn)
var ti = time.Now().In(time.UTC)
var createAt = timestamppb.New(ti)

func TestTodosServiceAdd(t *testing.T) {

	type args struct {
		ctx context.Context
		req *proto.AddRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test case #1",
			args{context.Background(),
				&proto.AddRequest{
					Api: apiVer,
					Todo: &proto.Todo{
						Id: 12,
						Title:       "Todo100",
						Description: "sdfsdf",
						CreateAt:    createAt,
						UpdateAt:    createAt,
						Status:      proto.Status_DONE,
						Tags:        []string{"Relax", "as", "Love"},
					},
				}},
			false,
		},
		// {"test case #2", args{context.Background(), &proto.AddRequest{
		// 	Api: "v1",
		// 	Todo: &proto.Todo{
		// 		Id:          1,
		// 		Title:       "asdasd",
		// 		Description: "asdasd",
		// 		CreateAt:    createAt,
		// 		UpdateAt:    createAt,
		// 		Status:      proto.Todo_TODO,
		// 		Tag: []string{"sao"},
		// 	},
		// }},
		// 	true,
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := client.Add(context.Background(), tt.args.req)
			fmt.Println(res)
			if (err != nil) != tt.wantErr {
				t.Errorf("TodosService.Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}

func TestTodosServiceGetOne(t *testing.T) {

	type args struct {
		ctx context.Context
		req *proto.GetOneRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"get todo exist in database",
			args{context.Background(),
				&proto.GetOneRequest{
					Api: apiVer,
					Id:  1,
				}},
			false,
		},
		{
			"get todo doesn't exist in database",
			args{context.Background(),
				&proto.GetOneRequest{
					Api: apiVer,
					Id:  10000,
				}},

			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := client.GetOne(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("TodosService.GetOne() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(got)
		})
	}
}

func TestTodosServiceGetAll(t *testing.T) {

	type args struct {
		ctx context.Context
		req *proto.GetAllRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"test case #1",
			args{context.Background(),
				&proto.GetAllRequest{
					Api: apiVer,
					Pag: 3,
					Status: []proto.Status{},
					Tags:   []string{"sleep"},
				}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := client.GetAll(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("TodosService.GetOne() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(got)
		})
	}
}

func TestTodosServiceDelete(t *testing.T) {

	type args struct {
		ctx context.Context
		req *proto.DeleteRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"delete record doesn't exist in database",
			args{context.Background(),
				&proto.DeleteRequest{
					Api: apiVer,
					Id:  10000,
				}},
			true,
		},
		{
			"delete record  exist in database",
			args{context.Background(),
				&proto.DeleteRequest{
					Api: apiVer,
					Id:  1,
				}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := client.Delete(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("TodosService.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}

func TestTodosServiceUpdate(t *testing.T) {

	type args struct {
		ctx context.Context
		req *proto.UpdateRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"updating record doesn't exist in database",
			args{context.Background(),
				&proto.UpdateRequest{
					Api: apiVer,
					Todo: &proto.Todo{
						Id:          1000,
						Title:       "asdasd",
						Description: "asdasd",
						CreateAt:    createAt,
						UpdateAt:    createAt,
						Status:      proto.Status_TODO,
					},
				}},
			true,
		},
		{
			"updating record exist in database",
			args{context.Background(),
				&proto.UpdateRequest{
					Api: apiVer,
					Todo: &proto.Todo{
						Id:          3,
						Title:       "duckhue01",
						Description: "duckhue01",
						CreateAt:    createAt,
						UpdateAt:    createAt,
						Status:      proto.Status_TODO,
					},
				}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			_, err := client.Update(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("TodosService.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestTodosServiceReorder(t *testing.T) {

	type args struct {
		ctx context.Context
		req *proto.ReorderRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// {
		// 	"reorder record doesn't exist in database",
		// 	args{context.Background(),
		// 		&proto.ReorderRequest{
		// 			Api: apiVer,
		// 			Id:  102,
		// 			Pos: 1,
		// 		}},
		// 	true,
		// },
		{
			"reorder record  exist in database ",
			args{context.Background(),
				&proto.ReorderRequest{
					Api: apiVer,
					Id:  2,
					Pos: 1,
				}},
			false,
		},
		// {
		// 	"reorder record  exist in database",
		// 	args{context.Background(),
		// 		&proto.ReorderRequest{
		// 			Api: apiVer,
		// 			Id:  1,
		// 			Pos: 3,
		// 		}},
		// 	false,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := client.Reorder(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("TodosService.Reorder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}
