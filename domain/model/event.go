package model

type Event struct {
	message string
}

func NewEvent(message string) Event {
	return Event{
		message: message,
	}
}

func (e Event) GetMessage() string {
	return e.message
}
