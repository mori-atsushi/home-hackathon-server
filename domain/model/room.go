package model

import (
	"sync"
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
	r.mutex.Unlock()
}

func (r Room) RemoveChannel(user User) {
	r.mutex.Lock()
	close(r.channels[user.id].channel)
	delete(r.channels, user.id)
	r.mutex.Unlock()
}

func (r Room) SendEvent(user User, event Event) {
	r.mutex.Lock()
	for _, channel := range r.channels {
		channel.SendEvent(event)
	}
	r.mutex.Unlock()
}

func (r Room) ReceiveEvent(user User) <-chan Event {
	r.mutex.Lock()
	channel := r.channels[user.id]
	r.mutex.Unlock()
	return channel.channel
}
