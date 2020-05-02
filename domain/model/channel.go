package model

type Channel struct {
	channel chan Event
}

func NewChannel() *Channel {
	return &Channel{
		channel: make(chan Event),
	}
}
