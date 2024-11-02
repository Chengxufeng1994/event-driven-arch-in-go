package out

import (
	"context"
)

type CustomerCacheRepository interface {
	Add(ctx context.Context, customerID, name string) error
	CustomerRepository
}
