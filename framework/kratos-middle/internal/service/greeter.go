package service

import (
	"context"

	v1 "kratos-middle/api/helloworld/v1"
	"kratos-middle/internal/biz"
)

// GreeterService is a greeter service.
type GreeterService struct {
	v1.UnimplementedGreeterServer

	uc *biz.GreeterUsecase
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase) *GreeterService {
	return &GreeterService{uc: uc}
}

// SayHello implements helloworld.GreeterServer.
func (s *GreeterService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	g, err := s.uc.CreateGreeter(ctx, &biz.Greeter{Hello: in.Name})
	if err != nil {
		return nil, err
	}
	return &v1.HelloReply{Message: "Hello " + g.Hello}, nil
}

func (s *GreeterService) GetData(ctx context.Context, req *v1.GetDataRequest) (*v1.GetDataResponse, error) {
	data, err := s.uc.GetData(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.GetDataResponse{Data: data}, nil
}
