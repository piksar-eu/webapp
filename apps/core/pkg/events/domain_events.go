package events

type DomainEvents struct {
	events []Event
}

func (e *DomainEvents) Record(ev Event) {
	e.events = append(e.events, ev)
}

func (e *DomainEvents) PullEvents() []Event {
	ev := e.events
	e.events = nil
	return ev
}
