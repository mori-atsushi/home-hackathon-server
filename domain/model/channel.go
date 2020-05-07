package model

const channelBufferSize = 10

type Channel struct {
	channel chan Event
}

func NewChannel() *Channel {
	return &Channel{
		channel: make(chan Event, channelBufferSize),
	}
}

func (c Channel) SendEvent(event Event) {
	go func() {
		c.channel <- event
	}()
}
