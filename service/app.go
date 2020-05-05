package service

import (
	"github.com/Mori-Atsushi/home-hackathon-server/domain/model"
	"github.com/Mori-Atsushi/home-hackathon-server/domain/usecase"
	"github.com/Mori-Atsushi/home-hackathon-server/pb"
)

type AppService struct {
	room model.Room
}

func (s *AppService) Event(srv pb.AppService_EventServer) error {
	user := model.NewUser()
	usecase.JoinRoom(&s.room, user)
	usecase.ObserveRoom(&s.room, user, srv)
	usecase.LeaveRoom(&s.room, user)
	return nil
}

func NewAppService() *AppService {
	return &AppService{
		room: model.NewRoom(),
	}
}
