package query

import "time"

type (
	Filters struct {
		CustomerID string
		After      time.Time
		Before     time.Time
		StoreIDs   []string
		ProductIDs []string
		MinTotal   float64
		MaxTotal   float64
		Status     string
	}

	SearchOrders struct {
		Filters Filters
		Next    string
		Limit   int
	}
)
