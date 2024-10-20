package valueobject

type OrderStatus string

const (
	OrderUnknown   OrderStatus = ""
	OrderPending   OrderStatus = "pending"
	OrderInProcess OrderStatus = "in-progress"
	OrderReady     OrderStatus = "ready"
	OrderCompleted OrderStatus = "completed"
	OrderCancelled OrderStatus = "canceled"
)

func NewOrderStatus(status string) OrderStatus {
	switch status {
	case OrderPending.String():
		return OrderPending
	case OrderInProcess.String():
		return OrderInProcess
	case OrderReady.String():
		return OrderReady
	case OrderCompleted.String():
		return OrderCompleted
	case OrderCancelled.String():
		return OrderCancelled
	default:
		return OrderUnknown
	}
}

func (s OrderStatus) String() string {
	switch s {
	case OrderPending, OrderInProcess, OrderReady, OrderCompleted, OrderCancelled:
		return string(s)
	default:
		return ""
	}
}
