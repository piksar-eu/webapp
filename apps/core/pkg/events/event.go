package events

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	MessageId string
	OccuredAt time.Time
	Type      string
	Payload   interface{}
}

func NewEvent(t string, p any) Event {
	return Event{
		Type:      t,
		Payload:   p,
		MessageId: uuid.NewString(),
		OccuredAt: time.Now().UTC(),
	}
}
