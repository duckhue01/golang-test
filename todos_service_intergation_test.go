package main

import (
	"context"
	"testing"
	"time"

	proto "github.com/duckhue01/golang_test/proto/v1"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

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
					Api: "v1",
					Todo: &proto.Todo{
						Id:          1,
						Title:       "asdasd",
						Description: "asdasd",
						CreateAt:    createAt,
						UpdateAt:    createAt,
						IsDone:      false,
					},
				}},
			false,
		},
		{"test case #2", args{context.Background(), &proto.AddRequest{
			Api: "v2",
			Todo: &proto.Todo{
				Id:          1,
				Title:       "asdasd",
				Description: "asdasd",
				CreateAt:    createAt,
				UpdateAt:    createAt,
				IsDone:      false,
			},
		}},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := client.Add(context.Background(), tt.args.req)
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
					Api: "v1",
					Id:  2,
				}},

			false,
		},
		{
			"get todo doesn't exist in database",
			args{context.Background(),
				&proto.GetOneRequest{
					Api: "v1",
					Id:  10000,
				}},

			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := client.GetOne(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("TodosService.GetOne() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
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
					Api: "v1",
				}},

			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := client.GetAll(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("TodosService.GetOne() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
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
					Api: "v1",
					Id:  10000,
				}},
			true,
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
					Api: "v1",
					Todo: &proto.Todo{
						Id:          1000,
						Title:       "asdasd",
						Description: "asdasd",
						CreateAt:    createAt,
						UpdateAt:    createAt,
						IsDone:      false,
					},
				}},
			true,
		},
		{
			"updating record exist in database",
			args{context.Background(),
				&proto.UpdateRequest{
					Api: "v1",
					Todo: &proto.Todo{
						Id:          3,
						Title:       "duckhue01",
						Description: "duckhue01",
						CreateAt:    createAt,
						UpdateAt:    createAt,
						IsDone:      true,
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
