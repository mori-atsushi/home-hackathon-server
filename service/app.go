package service

import (
	"context"
	"log"

	"github.com/Mori-Atsushi/home-hackathon-server/domain/model"
	"github.com/Mori-Atsushi/home-hackathon-server/pb"
)

type AppService struct {
	room model.Room
}

func (s *AppService) Ping(ctx context.Context, req *pb.PingRequest) (*pb.PingResponse, error) {
	log.Printf("Received: %v", req)
	return &pb.PingResponse{Message: "pong"}, nil
}

func (s *AppService) Event(srv pb.AppService_EventServer) error {
	user := model.NewUser()
	s.room.AddChannel(user)
	log.Printf("new: %v", user)
	for {
		resp, err := srv.Recv()
		if err != nil {
			defer s.room.RemoveChannel(user)
			log.Printf("close: %v", user)
			break
		}
		event := model.NewEvent(resp.GetEvent().GetMessage())
		s.room.SendEvent(user, event)
		log.Printf("recieve: %v", resp)
	}
	return nil
}

func NewAppService() *AppService {
	return &AppService{
		room: model.NewRoom(),
	}
}
