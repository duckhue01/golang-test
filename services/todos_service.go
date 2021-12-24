package services

import (
	"context"
	"fmt"

	proto "github.com/duckhue01/golang_test/proto/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TodosStore interface {
	Add(context.Context, *proto.AddRequest) error
	GetOne(context.Context, int32) (*proto.Todo, error)
	GetAll(ctx context.Context, req *proto.GetAllRequest) ([]*proto.Todo, error)
	Update(context.Context, *proto.UpdateRequest) error
	Delete(context.Context, int32) error
	Reorder(context.Context, *proto.ReorderRequest) error
}

type TodosService struct {
	Store TodosStore
}

const apiVer = "v2"

func NewTodosService(store TodosStore) *TodosService {
	return &TodosService{Store: store}
}

// checkAPI checks if the API version requested by client is supported by server
func (s *TodosService) checkAPI(api string) error {
	// API version is "" means use current version of the service
	if len(api) > 0 {
		if apiVer != api {
			return status.Errorf(codes.Unimplemented,
				"unsupported API version: service implements API version '%s', but asked for '%s'", apiVer, api)
		}
	}
	return nil
}

func (s *TodosService) Add(ctx context.Context, req *proto.AddRequest) (*proto.AddResponse, error) {
	if err := s.checkAPI(req.Api); err != nil {
		return &proto.AddResponse{
			Api: apiVer,
		}, err
	}

	fmt.Println(req.Todo.Tags)
	err := s.Store.Add(ctx, req)
	if err != nil {
		return &proto.AddResponse{
			Api: apiVer,
		}, err
	}

	return &proto.AddResponse{
		Api: apiVer,
	}, nil
}

func (s *TodosService) GetOne(ctx context.Context, req *proto.GetOneRequest) (*proto.GetOneResponse, error) {
	if err := s.checkAPI(req.Api); err != nil {
		return &proto.GetOneResponse{
			Api: apiVer,
		}, err
	}

	todo, err := s.Store.GetOne(ctx, req.Id)
	if err != nil {
		return &proto.GetOneResponse{
			Api: apiVer,
		}, err
	}

	return &proto.GetOneResponse{
		Api:  apiVer,
		Todo: todo,
	}, nil
}

func (s *TodosService) GetAll(ctx context.Context, req *proto.GetAllRequest) (*proto.GetAllResponse, error) {
	if err := s.checkAPI(req.Api); err != nil {
		return &proto.GetAllResponse{
			Api: apiVer,
		}, err
	}
	fmt.Println(req)

	todos, err := s.Store.GetAll(ctx, req)
	if err != nil {
		return &proto.GetAllResponse{
			Api: apiVer,
		}, err
	}

	return &proto.GetAllResponse{
		Api:  apiVer,
		Todo: todos,
	}, nil
}

func (s *TodosService) Update(ctx context.Context, req *proto.UpdateRequest) (*proto.UpdateResponse, error) {
	if err := s.checkAPI(req.Api); err != nil {
		return &proto.UpdateResponse{
			Api: apiVer,
		}, err
	}

	err := s.Store.Update(ctx, req)
	if err != nil {
		return &proto.UpdateResponse{
			Api: apiVer,
		}, err
	}

	return &proto.UpdateResponse{
		Api: apiVer,
	}, nil
}

func (s *TodosService) Delete(ctx context.Context, req *proto.DeleteRequest) (*proto.DeleteResponse, error) {
	if err := s.checkAPI(req.Api); err != nil {
		return &proto.DeleteResponse{
			Api: apiVer,
		}, err
	}

	err := s.Store.Delete(ctx, req.Id)
	if err != nil {
		return &proto.DeleteResponse{
			Api: apiVer,
		}, err
	}

	return &proto.DeleteResponse{
		Api: apiVer,
	}, nil
}

func (s *TodosService) Reorder(ctx context.Context, req *proto.ReorderRequest) (*proto.ReorderResponse, error) {
	if err := s.checkAPI(req.Api); err != nil {
		return &proto.ReorderResponse{
			Api: apiVer,
		}, err
	}

	err := s.Store.Reorder(ctx, req)
	if err != nil {
		return &proto.ReorderResponse{
			Api: apiVer,
		}, err
	}

	return &proto.ReorderResponse{
		Api: apiVer,
	}, nil
}

// i have no idea why this doesn't work
// so i decide to comment this method at proto/v2/todos_grpc.pb.go
// func (s *TodosService) mustEmbedUnimplementedTodosServiceTodosService() {
// }
