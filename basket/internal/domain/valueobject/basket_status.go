package valueobject

type BasketStatus string

const (
	BasketUnknown    BasketStatus = ""
	BasketOpen       BasketStatus = "open"
	BasketCancelled  BasketStatus = "canceled"
	BasketCheckedOut BasketStatus = "checked_out"
)

func NewBasketStatus(status string) (BasketStatus, error) {
	bs := BasketStatus(status)

	return bs, nil
}

func (status BasketStatus) IsValid() bool {
	switch status {
	case BasketOpen, BasketCancelled, BasketCheckedOut:
		return true
	default:
		return false
	}
}

func (status BasketStatus) String() string {
	switch status {
	case BasketOpen, BasketCancelled, BasketCheckedOut:
		return string(status)
	default:
		return ""
	}
}
