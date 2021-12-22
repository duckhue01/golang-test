package services

import (
	"context"
	"fmt"
	"reflect"
	"testing"
	"time"

	proto "github.com/duckhue01/golang_test/proto/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type InMemoStore struct {
	todos []proto.Todo
}

func (i *InMemoStore) Add(ctx context.Context, req *proto.AddRequest) error {

	i.todos = append(i.todos, *req.GetTodo())
	return nil

}
func (i *InMemoStore) GetOne(ctx context.Context, id int32) (*proto.Todo, error) {

	for idx := 0; idx < len(i.todos); idx++ {
		if i.todos[idx].Id == id {
			return &i.todos[idx], nil
		}
	}

	return nil, fmt.Errorf("Todo with ID='%d' is not found", id)
}

func (i *InMemoStore) GetAll(ctx context.Context) ([]*proto.Todo, error) {
	var todos []*proto.Todo
	for idx := 0; idx < len(i.todos); idx++ {
		todos = append(todos, &i.todos[idx])
	}

	return todos, nil

}

func (i *InMemoStore) Update(ctx context.Context, req *proto.UpdateRequest) error {

	for idx := 0; idx < len(i.todos); idx++ {
		if i.todos[idx].Id == req.Todo.Id {
			i.todos[idx] = *req.GetTodo()
			return nil
		}
	}

	return fmt.Errorf("Todo with ID='%d' is not found", req.Todo.Id)
}

func (i *InMemoStore) Delete(ctx context.Context, id int32) error {
	for idx1 := 0; idx1 < len(i.todos); idx1++ {
		if i.todos[idx1].Id == id {
			for idx2 := idx1 + 1; idx2 < len(i.todos); idx2++ {
				i.todos[idx1] = i.todos[idx2]
				idx1++
			}
			return nil
		}
	}
	return fmt.Errorf("Todo with ID='%d' is not found", id)
}

var s = NewTodosService(&InMemoStore{})

func TestTodosService_checkAPI(t *testing.T) {

	type args struct {
		api string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test v1", args{api: "v1"}, false},
		{"test v2", args{api: "v2"}, true},
		{"test v3", args{api: "v3"}, true},
		{"test v4", args{api: "v4"}, true},
		{"test -1", args{api: "-1"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if err := s.checkAPI(tt.args.api); (err != nil) != tt.wantErr {
				t.Errorf("TodosService.checkAPI() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTodosService_Add(t *testing.T) {
	ti := time.Now().In(time.UTC)
	createAt := timestamppb.New(ti)

	type args struct {
		ctx context.Context
		req *proto.AddRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *proto.AddResponse
		wantErr bool
	}{
		{"test case #1", args{context.Background(), &proto.AddRequest{
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
			&proto.AddResponse{
				Api:  apiVer,
			}, false,
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
			&proto.AddResponse{
				Api:  apiVer,
			}, true,
		},
	}


	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := s.Add(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("TodosService.Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TodosService.Add() = %v, want %v", got, tt.want)
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
		want    *proto.GetOneResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := s.GetOne(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("TodosService.GetOne() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TodosService.GetOne() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTodosServiceGetAll(t *testing.T) {
	type fields struct {
		Store TodosStore
	}
	type args struct {
		ctx context.Context
		req *proto.GetAllRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *proto.GetAllResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := s.GetAll(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("TodosService.GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TodosService.GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTodosServiceUpdate(t *testing.T) {
	type fields struct {
		Store TodosStore
	}
	type args struct {
		ctx context.Context
		req *proto.UpdateRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *proto.UpdateResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := s.Update(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("TodosService.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TodosService.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTodosServiceDelete(t *testing.T) {
	type fields struct {
		Store TodosStore
	}
	type args struct {
		ctx context.Context
		req *proto.DeleteRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *proto.DeleteResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := s.Delete(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("TodosService.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TodosService.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}


