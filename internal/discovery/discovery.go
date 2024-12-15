package discovery

import (
	"context"
)

type Discovery interface {
	Register(ctx context.Context, serviceName, instanceID, host string, port int) error
	DeRegister(ctx context.Context, serviceName, instanceID string) error
}
