package services

import (
	"context"
	"fmt"
	"testing"
	"time"

	proto "github.com/duckhue01/golang_test/proto/v2"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var ti = time.Now().In(time.UTC)
var createAt = timestamppb.New(ti)

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

func (i *InMemoStore) GetAll(ctx context.Context, req *proto.GetAllRequest) ([]*proto.Todo, error) {
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

func (i *InMemoStore) Reorder(ctx context.Context, req *proto.ReorderRequest) error {
	// for idx1 := 0; idx1 < len(i.todos); idx1++ {
	// 	if i.todos[idx1].Id == id {
	// 		for idx2 := idx1 + 1; idx2 < len(i.todos); idx2++ {
	// 			i.todos[idx1] = i.todos[idx2]
	// 			idx1++
	// 		}
	// 		return nil
	// 	}
	// }
	// return fmt.Errorf("Todo with ID='%d' is not found", id)

	return nil
}

var s = NewTodosService(&InMemoStore{
	todos: []proto.Todo{
		{
			Id:          1,
			Title:       "duckhue02",
			Description: "sdfsdf",
			CreateAt:    createAt,
			UpdateAt:    createAt,
			Status:      proto.Status_DOING,
			Tags:        []string{"Relax", "Weekly", "Love"},
		},
		{
			Id:          2,
			Title:       "duckhue02",
			Description: "sdfsdf",
			CreateAt:    createAt,
			UpdateAt:    createAt,
			Status:      proto.Status_DOING,
			Tags:        []string{"Relax", "Weekly", "Love"},
		},
		{
			Id:          3,
			Title:       "duckhue02",
			Description: "sdfsdf",
			CreateAt:    createAt,
			UpdateAt:    createAt,
			Status:      proto.Status_DOING,
			Tags:        []string{"Relax", "Weekly", "Love"},
		},

		{
			Id:          4,
			Title:       "duckhue02",
			Description: "sdfsdf",
			CreateAt:    createAt,
			UpdateAt:    createAt,
			Status:      proto.Status_DOING,
			Tags:        []string{"Relax", "Weekly", "Love"},
		},
		{
			Id:          5,
			Title:       "duckhue02",
			Description: "sdfsdf",
			CreateAt:    createAt,
			UpdateAt:    createAt,
			Status:      proto.Status_DOING,
			Tags:        []string{"Relax", "Weekly", "Love"},
		},
		{
			Id:          6,
			Title:       "duckhue02",
			Description: "sdfsdf",
			CreateAt:    createAt,
			UpdateAt:    createAt,
			Status:      proto.Status_DOING,
			Tags:        []string{"Relax", "Weekly", "Love"},
		},
	},
})

func TestTodosServiceAdd(t *testing.T) {
	ti := time.Now().In(time.UTC)
	createAt := timestamppb.New(ti)

	type args struct {
		ctx context.Context
		req *proto.AddRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test if todo id added successfully", args{context.Background(), &proto.AddRequest{
			Api: "v2",
			Todo: &proto.Todo{
				Title:       "Todo1",
				Description: "asdasd",
				CreateAt:    createAt,
				UpdateAt:    createAt,
				Status:      proto.Status_TODO,
			},
		}}, false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			_, err := s.Add(tt.args.ctx, tt.args.req)
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
					Api: "v2",
					Id:  2,
				}},

			false,
		},
		{
			"get todo doesn't exist in database",
			args{context.Background(),
				&proto.GetOneRequest{
					Api: "v2",
					Id:  10000,
				}},

			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			_, err := s.GetOne(tt.args.ctx, tt.args.req)
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
					Api: "v2",
				}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			_, err := s.GetAll(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("TodosService.GetAll() error = %v, wantErr %v", err, tt.wantErr)
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
					Api: "v2",
					Todo: &proto.Todo{
						Id:          1000,
						Title:       "asdasd",
						Description: "asdasd",
						CreateAt: &timestamp.Timestamp{
							Seconds: 123123123,
							Nanos:   123123123,
						},
						UpdateAt: &timestamp.Timestamp{
							Seconds: 123123123,
							Nanos:   123123123,
						},
						Status: proto.Status_DOING,
					},
				}},
			true,
		},
		{
			"updating record exist in database",
			args{context.Background(),
				&proto.UpdateRequest{
					Api: "v2",
					Todo: &proto.Todo{
						Id:          3,
						Title:       "duckhue01",
						Description: "duckhue01",
						CreateAt: &timestamp.Timestamp{
							Seconds: 123123123,
							Nanos:   123123123,
						},
						UpdateAt: &timestamp.Timestamp{
							Seconds: 123123123,
							Nanos:   123123123,
						},
						Status: proto.Status_DOING,
					},
				}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			_, err := s.Update(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("TodosService.Update() error = %v, wantErr %v", err, tt.wantErr)
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
		name string

		args    args
		wantErr bool
	}{
		{
			"delete record doesn't exist in database",
			args{context.Background(),
				&proto.DeleteRequest{
					Api: "v2",
					Id:  10000,
				}},
			true,
		},
		{
			"delete record exist in database",
			args{context.Background(),
				&proto.DeleteRequest{
					Api: "v2",
					Id:  1,
				}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			_, err := s.Delete(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("TodosService.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}
