package services

import (
	"context"

	proto "github.com/duckhue01/golang_test/proto/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TodosStore interface {
	Add(context.Context, *proto.AddRequest) error
	GetOne(context.Context, int32) (*proto.Todo, error)
	GetAll(context.Context) ([]*proto.Todo, error)
	Update(context.Context, *proto.UpdateRequest) error
	Delete(context.Context, int32) error
}

type TodosService struct {
	Store TodosStore
}

const apiVer = "v1"

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
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return &proto.AddResponse{
			Api:  apiVer,
			IsOk: false,
			Mes:  err.Error(),
		}, err
	}

	err := s.Store.Add(ctx, req)

	if err != nil {
		return &proto.AddResponse{
			Api:  apiVer,
			IsOk: false,
			Mes:  err.Error(),
		}, err

	}
	return &proto.AddResponse{
		Api:  apiVer,
		IsOk: true,
	}, nil
}

func (s *TodosService) GetOne(ctx context.Context, req *proto.GetOneRequest) (*proto.GetOneResponse, error) {
	if err := s.checkAPI(req.Api); err != nil {
		return &proto.GetOneResponse{
			Api:  apiVer,
			IsOk: false,
			Mes:  err.Error(),
		}, err
	}

	todo, err := s.Store.GetOne(ctx, req.Id)
	if err != nil {
		return &proto.GetOneResponse{
			Api:  apiVer,
			IsOk: false,
			Mes:  err.Error(),
		}, err
	}

	return &proto.GetOneResponse{
		Api:  apiVer,
		IsOk: true,
		Todo: todo,
	}, nil
}

func (s *TodosService) GetAll(ctx context.Context, req *proto.GetAllRequest) (*proto.GetAllResponse, error) {

	if err := s.checkAPI(req.Api); err != nil {
		return &proto.GetAllResponse{
			Api:  apiVer,
			IsOk: false,
			Mes:  err.Error(),
		}, err
	}

	todo, err := s.Store.GetAll(ctx)
	if err != nil {
		return &proto.GetAllResponse{
			Api:  apiVer,
			IsOk: false,
			Mes:  err.Error(),
		}, err
	}

	return &proto.GetAllResponse{
		Api:  apiVer,
		IsOk: true,
		Todo: todo,
	}, nil
}

func (s *TodosService) Update(ctx context.Context, req *proto.UpdateRequest) (*proto.UpdateResponse, error) {
	if err := s.checkAPI(req.Api); err != nil {
		return &proto.UpdateResponse{
			Api:  apiVer,
			IsOk: false,
			Mes:  err.Error(),
		}, err
	}

	err := s.Store.Update(ctx, req)
	if err != nil {
		return &proto.UpdateResponse{
			Api:  apiVer,
			IsOk: false,
			Mes:  err.Error(),
		}, err
	}

	return &proto.UpdateResponse{
		Api:  apiVer,
		IsOk: true,
	}, nil
}

func (s *TodosService) Delete(ctx context.Context, req *proto.DeleteRequest) (*proto.DeleteResponse, error) {
	if err := s.checkAPI(req.Api); err != nil {
		return &proto.DeleteResponse{
			Api:  apiVer,
			IsOk: false,
			Mes:  err.Error(),
		}, err
	}

	err := s.Store.Delete(ctx, req.Id)
	if err != nil {
		return &proto.DeleteResponse{
			Api:  apiVer,
			IsOk: false,
			Mes:  err.Error(),
		}, err
	}

	return &proto.DeleteResponse{
		Api:  apiVer,
		IsOk: true,
	}, nil
}

// i have no idea why this doesn't work
// func (s *TodosService) MustEmbedUnimplementedTodosServiceTodosService() {

// }
