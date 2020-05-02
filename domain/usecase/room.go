package usecase

import (
	"log"

	"github.com/Mori-Atsushi/home-hackathon-server/domain/model"
	"github.com/Mori-Atsushi/home-hackathon-server/pb"
)

func JoinRoom(room *model.Room, user model.User) {
	room.AddChannel(user)
	log.Printf("new: %v", user)
}

func ObserveRoom(room *model.Room, user model.User, srv pb.AppService_EventServer) {
	done := make(chan int)
	go func() {
		sendEvent(room, user, srv)
		done <- 1
	}()
	go func() {
		receiveEvent(room, user, srv)
		done <- 1
	}()
	// 片方が呼ばれたら終了
	<-done
}

func LeaveRoom(room *model.Room, user model.User) {
	room.RemoveChannel(user)
	log.Printf("remove: %v", user)
}

func sendEvent(room *model.Room, user model.User, srv pb.AppService_EventServer) {
	for {
		req, err := srv.Recv()
		if err != nil {
			log.Printf("close: %v", user)
			break
		}
		event := model.NewEvent(req.GetEvent())
		room.SendEvent(user, event)
		log.Printf("recieve: %v, %v", user, req)
	}
}

func receiveEvent(room *model.Room, user model.User, srv pb.AppService_EventServer) {
	channel := room.ReceiveEvent(user)
	for {
		event, ok := <-channel
		if !ok {
			break
		}
		resp := pb.EventResponse{Event: event.GetRaw()}
		log.Printf("send: %v, %v", user, resp)
		srv.Send(&resp)
	}
}
