package model

import (
	"sync"

	"github.com/Mori-Atsushi/home-hackathon-server/pb"
)

type Room struct {
	mutex    *sync.Mutex
	channels map[string]*Channel
}

func NewRoom() Room {
	return Room{
		mutex:    &sync.Mutex{},
		channels: map[string]*Channel{},
	}
}

func (r Room) AddChannel(user User) {
	channel := NewChannel()
	r.mutex.Lock()
	r.channels[user.id] = channel
	r.sendUserEvent()
	r.mutex.Unlock()
}

func (r Room) RemoveChannel(user User) {
	r.mutex.Lock()
	close(r.channels[user.id].channel)
	delete(r.channels, user.id)
	r.sendUserEvent()
	r.mutex.Unlock()
}

func (r Room) SendSoundEvent(user User, sound *pb.Sound) {
	event := NewSoundEvent(user.id, sound)
	r.mutex.Lock()
	r.sendEvent(event)
	r.mutex.Unlock()
}

func (r Room) sendUserEvent() {
	keys := []string{}
	for key := range r.channels {
		keys = append(keys, key)
	}
	event := NewUsersEvent(keys)
	r.sendEvent(event)
}

func (r Room) sendEvent(event Event) {
	for _, channel := range r.channels {
		channel.SendEvent(event)
	}
}

func (r Room) ReceiveEvent(user User) <-chan Event {
	r.mutex.Lock()
	channel := r.channels[user.id]
	r.mutex.Unlock()
	return channel.channel
}
