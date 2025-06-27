package command

import (
	"context"
	"fmt"
	"log"
	"sync"
)

// CommandBus - główny dispatcher
type CommandBus struct {
	handlers map[string]CommandHandler
	mu       sync.RWMutex
}

func NewCommandBus() *CommandBus {
	return &CommandBus{
		handlers: make(map[string]CommandHandler),
	}
}

func (cb *CommandBus) RegisterHandler(command Command, handler CommandHandler) {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	cb.handlers[QualifiedType(command)] = handler
}

func (cb *CommandBus) Dispatch(ctx context.Context, cmd Command) (any, error) {
	switch cmd.GetMode() {
	case CommandAsync:
		payload, err := cmd.Marshal()
		if err != nil {
			return nil, err
		}
		// return cb.queue.Publish(payload) // pseudo: RabbitMQ publish
		log.Println("async command payload", payload)
		return nil, nil
	default:
		handler, ok := cb.handlers[QualifiedType(cmd)]
		if !ok {
			return nil, fmt.Errorf("no handler registered for type %s", QualifiedType(cmd))
		}

		return handler.Handle(ctx, cmd)
	}
}
