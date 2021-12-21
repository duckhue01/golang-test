package services

import (
	"context"
	"reflect"
	"testing"

	proto "github.com/duckhue01/golang_test/proto/v1"
)



type todosStore struct {
	todos []proto.AddRequest
}

func (s *todosStore) Add(req *proto.AddRequest) bool {
	
	for i := range s.todos {
		if i == int(req.Todo.Id) {
			return false
		}
	}

	s.todos = append(s.todos, proto.AddRequest{})
	return true
}

var s = NewTodosService(&todosStore{})

func TestTodosService_Add(t *testing.T) {
	type args struct {
		ctx context.Context
		req *proto.AddRequest
	}
	tests := []struct {
		name    string
		s       *TodosService
		args    args
		want    *proto.AddResponse
		wantErr bool
	}{
		// TODO: Add test cases.
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

// // func TestTodosService_GetOne(t *testing.T) {
// // 	type args struct {
// // 		ctx context.Context
// // 		req *proto.GetOneRequest
// // 	}
// // 	tests := []struct {
// // 		name    string
// // 		s       *TodosService
// // 		args    args
// // 		want    *proto.GetOneResponse
// // 		wantErr bool
// // 	}{
// // 		// TODO: Add test cases.
// // 	}
// // 	for _, tt := range tests {
// // 		t.Run(tt.name, func(t *testing.T) {
// // 			got, err := s.GetOne(tt.args.ctx, tt.args.req)
// // 			if (err != nil) != tt.wantErr {
// // 				t.Errorf("TodosService.GetOne() error = %v, wantErr %v", err, tt.wantErr)
// // 				return
// // 			}
// // 			if !reflect.DeepEqual(got, tt.want) {
// // 				t.Errorf("TodosService.GetOne() = %v, want %v", got, tt.want)
// // 			}
// // 		})
// // 	}
// // }

// // func TestTodosService_GetAll(t *testing.T) {
// // 	type args struct {
// // 		req *proto.GetAllRequest
// // 		rvc proto.TodosService_GetAllServer
// // 	}
// // 	tests := []struct {
// // 		name    string
// // 		s       *TodosService
// // 		args    args
// // 		wantErr bool
// // 	}{
// // 		// TODO: Add test cases.
// // 	}
// // 	for _, tt := range tests {
// // 		t.Run(tt.name, func(t *testing.T) {
// // 			if err := s.GetAll(tt.args.req, tt.args.rvc); (err != nil) != tt.wantErr {
// // 				t.Errorf("TodosService.GetAll() error = %v, wantErr %v", err, tt.wantErr)
// // 			}
// // 		})
// // 	}
// // }

// // func TestTodosService_Update(t *testing.T) {
// // 	type args struct {
// // 		ctx context.Context
// // 		req *proto.UpdateRequest
// // 	}
// // 	tests := []struct {
// // 		name    string
// // 		s       *TodosService
// // 		args    args
// // 		want    *proto.UpdateResponse
// // 		wantErr bool
// // 	}{
// // 		// TODO: Add test cases.
// // 	}
// // 	for _, tt := range tests {
// // 		t.Run(tt.name, func(t *testing.T) {

// // 			got, err := s.Update(tt.args.ctx, tt.args.req)
// // 			if (err != nil) != tt.wantErr {
// // 				t.Errorf("TodosService.Update() error = %v, wantErr %v", err, tt.wantErr)
// // 				return
// // 			}
// // 			if !reflect.DeepEqual(got, tt.want) {
// // 				t.Errorf("TodosService.Update() = %v, want %v", got, tt.want)
// // 			}
// // 		})
// // 	}
// // }

// // func TestTodosService_Delete(t *testing.T) {
// // 	type args struct {
// // 		ctx context.Context
// // 		req *proto.DeleteRequest
// // 	}
// // 	tests := []struct {
// // 		name    string
// // 		s       *TodosService
// // 		args    args
// // 		want    *proto.DeleteResponse
// // 		wantErr bool
// // 	}{
// // 		// TODO: Add test cases.
// // 	}
// // 	for _, tt := range tests {
// // 		t.Run(tt.name, func(t *testing.T) {

// // 			got, err := s.Delete(tt.args.ctx, tt.args.req)
// // 			if (err != nil) != tt.wantErr {
// // 				t.Errorf("TodosService.Delete() error = %v, wantErr %v", err, tt.wantErr)
// // 				return
// // 			}
// // 			if !reflect.DeepEqual(got, tt.want) {
// // 				t.Errorf("TodosService.Delete() = %v, want %v", got, tt.want)
// // 			}
// // 		})
// // 	}
// // }
