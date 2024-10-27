package entity

import "github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/domain/valueobject"

type Bot struct {
	ID        string
	Name      string
	BotStatus valueobject.BotStatus
}

func NewBot(id, name string, status valueobject.BotStatus) *Bot {
	return &Bot{
		ID:        id,
		Name:      name,
		BotStatus: status,
	}
}
