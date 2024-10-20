package command

type OrderCanceled struct {
	OrderID    string
	CustomerID string
}

func NewOrderCanceled(orderID, customerID string) OrderCanceled {
	return OrderCanceled{
		OrderID:    orderID,
		CustomerID: customerID,
	}
}
