package service

import (
	"context"
	"log"

	"github.com/Mori-Atsushi/home-hackathon-server/pb"
)

type AppService struct{}

func (s *AppService) Ping(ctx context.Context, req *pb.PingRequest) (*pb.PingResponse, error) {
	log.Printf("Received: %v", req)
	return &pb.PingResponse{Message: "pong"}, nil
}

func NewAppService() *AppService {
	return &AppService{}
}
