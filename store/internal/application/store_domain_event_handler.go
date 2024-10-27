package application

import "github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/application/port/in/event"

type (
	StoreDomainEventHandler struct {
		event.IgnoreUnimplementedDomainEventHandler
	}
)

var _ event.DomainEventHandlers = (*StoreDomainEventHandler)(nil)

func NewStoreDomainEventHandler() *StoreDomainEventHandler {
	return &StoreDomainEventHandler{}
}
