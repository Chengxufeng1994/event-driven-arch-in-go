package command

type OrderReady struct {
	OrderID    string
	CustomerID string
}

func NewOrderReady(orderID, customerID string) OrderReady {
	return OrderReady{
		OrderID:    orderID,
		CustomerID: customerID,
	}
}
