package valueobject

type OrderStatus string

const (
	OrderUnknown     OrderStatus = ""
	OrderIsPending   OrderStatus = "pending"
	OrderInProcess   OrderStatus = "in-progress"
	OrderIsReady     OrderStatus = "ready"
	OrderIsCompleted OrderStatus = "completed"
	OrderIsCancelled OrderStatus = "canceled"
)

func NewOrderStatus(status string) OrderStatus {
	switch status {
	case OrderIsPending.String():
		return OrderIsPending
	case OrderInProcess.String():
		return OrderInProcess
	case OrderIsReady.String():
		return OrderIsReady
	case OrderIsCompleted.String():
		return OrderIsCompleted
	case OrderIsCancelled.String():
		return OrderIsCancelled
	default:
		return OrderUnknown
	}
}

func (s OrderStatus) String() string {
	switch s {
	case OrderIsPending, OrderInProcess, OrderIsReady, OrderIsCompleted, OrderIsCancelled:
		return string(s)
	default:
		return ""
	}
}
