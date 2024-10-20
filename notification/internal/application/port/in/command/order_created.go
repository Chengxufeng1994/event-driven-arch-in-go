package command

type OrderCreated struct {
	OrderID    string
	CustomerID string
}

func NewOrderCreated(orderID, customerID string) OrderCreated {
	return OrderCreated{
		OrderID:    orderID,
		CustomerID: customerID,
	}
}
