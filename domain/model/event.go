package model

import (
	"github.com/Mori-Atsushi/home-hackathon-server/pb"
)

type Event struct {
	raw *pb.Event
}

func NewEvent(raw *pb.Event) Event {
	return Event{
		raw: raw,
	}
}

func (e Event) GetRaw() *pb.Event {
	return e.raw
}
