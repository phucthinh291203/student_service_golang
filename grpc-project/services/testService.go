package services

import (
	"context"
	pb "grpc-project/proto"
)

type TestService struct {
	pb.UnimplementedHelloServiceServer
}

// Hàm khởi tạo TestService
func NewTestService() *TestService {
	return &TestService{}
}

func (s *TestService) SendHelloMessage(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: "Hello " + req.Name}, nil
}
