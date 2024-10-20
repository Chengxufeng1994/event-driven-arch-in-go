package valueobject

type BotStatus string

const (
	BotUnknown BotStatus = ""
	BotIdle    BotStatus = "idle"
	BotActive  BotStatus = "active"
)

func NewBotStatus(status string) BotStatus {
	switch status {
	case BotIdle.String():
		return BotIdle
	case BotActive.String():
		return BotActive
	default:
		return BotUnknown
	}
}

func (s BotStatus) String() string {
	switch s {
	case BotIdle, BotActive:
		return string(s)
	default:
		return ""
	}
}
