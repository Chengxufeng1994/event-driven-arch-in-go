package aggregate

import "github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/domain/valueobject"

type BotAgg struct {
	ID        string
	Name      string
	BotStatus valueobject.BotStatus
}

func NewBot(id, name string, status valueobject.BotStatus) *BotAgg {
	return &BotAgg{
		ID:        id,
		Name:      name,
		BotStatus: status,
	}
}
