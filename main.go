package main

import (
	"log"
	"net"

	"github.com/Mori-Atsushi/home-hackathon-server/pb"
	"github.com/Mori-Atsushi/home-hackathon-server/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":8080"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAppServiceServer(s, service.NewAppService())
	reflection.Register(s)
	log.Printf("Listening on %v", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
