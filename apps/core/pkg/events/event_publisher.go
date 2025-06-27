package events

type EventPublisher interface {
	Publish(...Event)
	Register(string, EventListener)
}

type EventListener func(Event)
