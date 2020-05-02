package usecase

import (
	"log"
	"sync"

	"github.com/Mori-Atsushi/home-hackathon-server/domain/model"
	"github.com/Mori-Atsushi/home-hackathon-server/pb"
)

func JoinRoom(room *model.Room, user model.User) {
	room.AddChannel(user)
	log.Printf("new: %v", user)
}

func ObserveRoom(room *model.Room, user model.User, srv pb.AppService_EventServer) {
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		sendEvent(room, user, srv)
		wg.Done()
	}()
	go func() {
		receiveEvent(room, user, srv)
		wg.Done()
	}()
	wg.Wait()
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
		event := <-channel
		resp := pb.EventResponse{Event: event.GetRaw()}
		log.Printf("send: %v, %v", user, resp)
		srv.Send(&resp)
	}
}
