package valueobject

type BasketStatus string

const (
	BasketUnknown      BasketStatus = ""
	BasketIsOpen       BasketStatus = "open"
	BasketIsCanceled   BasketStatus = "canceled"
	BasketIsCheckedOut BasketStatus = "checked_out"
)

func NewBasketStatus(status string) (BasketStatus, error) {
	bs := BasketStatus(status)

	return bs, nil
}

func (status BasketStatus) IsValid() bool {
	switch status {
	case BasketIsOpen, BasketIsCanceled, BasketIsCheckedOut:
		return true
	default:
		return false
	}
}

func (status BasketStatus) String() string {
	switch status {
	case BasketIsOpen, BasketIsCanceled, BasketIsCheckedOut:
		return string(status)
	default:
		return ""
	}
}
