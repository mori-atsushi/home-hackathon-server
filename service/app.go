package service

import (
	"context"
	"log"
	"sync"

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
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		for {
			req, err := srv.Recv()
			if err != nil {
				log.Printf("close: %v", user)
				break
			}
			event := model.NewEvent(req.GetEvent().GetMessage())
			s.room.SendEvent(user, event)
			log.Printf("recieve: %v, %v", user, req)
		}
		wg.Done()
	}()
	go func() {
		channel := s.room.ReceiveEvent(user)
		for {
			event := <-channel
			resp := pb.EventResponse{Event: &pb.Event{Message: event.GetMessage()}}
			log.Printf("send: %v, %v", user, resp)
			srv.Send(&resp)
		}
		wg.Done()
	}()
	wg.Wait()
	defer s.room.RemoveChannel(user)
	return nil
}

func NewAppService() *AppService {
	return &AppService{
		room: model.NewRoom(),
	}
}
