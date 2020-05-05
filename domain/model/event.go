package model

import (
	"github.com/Mori-Atsushi/home-hackathon-server/pb"
)

type Event struct {
	raw *pb.EventResponse
}

func NewSoundEvent(userID string, sound *pb.Sound) Event {
	soundEvent := &pb.SoundEvent{
		Sound:  sound,
		UserID: userID,
	}
	raw := &pb.EventResponse{
		EventOneof: &pb.EventResponse_SoundEvent{
			SoundEvent: soundEvent,
		},
	}
	return Event{
		raw: raw,
	}
}

func NewUsersEvent(userIDs []string) Event {
	userEvent := &pb.UserEvent{
		UserIDs: userIDs,
	}
	raw := &pb.EventResponse{
		EventOneof: &pb.EventResponse_UserEvent{
			UserEvent: userEvent,
		},
	}
	return Event{
		raw: raw,
	}
}

func (e Event) GetRaw() *pb.EventResponse {
	return e.raw
}
