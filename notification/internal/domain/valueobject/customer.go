package valueobject

type Customer struct {
	ID        string
	Name      string
	SmsNumber string
}

func NewCustomer(id, name, smsNumber string) Customer {
	return Customer{
		ID:        id,
		Name:      name,
		SmsNumber: smsNumber,
	}
}
