package aggregate

type PaymentAgg struct {
	ID         string
	CustomerID string
	Amount     float64
}

func NewPayment(id, customerID string, amount float64) *PaymentAgg {
	return &PaymentAgg{
		ID:         id,
		CustomerID: customerID,
		Amount:     amount,
	}
}
