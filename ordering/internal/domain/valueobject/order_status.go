package valueobject

type OrderStatus string

const (
	OrderUnknown     OrderStatus = ""
	OrderIsPending   OrderStatus = "pending"
	OrderIsRejected  OrderStatus = "rejected"
	OrderIsApproved  OrderStatus = "approved"
	OrderIsInProcess OrderStatus = "in-progress"
	OrderIsReady     OrderStatus = "ready"
	OrderIsCompleted OrderStatus = "completed"
	OrderIsCancelled OrderStatus = "cancelled"
)

func NewOrderStatus(status string) OrderStatus {
	switch status {
	case OrderIsPending.String():
		return OrderIsPending
	case OrderIsRejected.String():
		return OrderIsRejected
	case OrderIsApproved.String():
		return OrderIsApproved
	case OrderIsInProcess.String():
		return OrderIsInProcess
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
	case OrderIsPending, OrderIsRejected, OrderIsApproved, OrderIsInProcess, OrderIsReady, OrderIsCompleted, OrderIsCancelled:
		return string(s)
	default:
		return ""
	}
}
