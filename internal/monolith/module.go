package monolith

import "context"

type Module interface {
	Name() string
	PrepareRun(ctx context.Context, mono Monolith) error
}
