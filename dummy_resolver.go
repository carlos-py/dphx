package dphx

import (
	"net"

	"golang.org/x/net/context"
)

// DummyResolver stores unresolved host address at context for later use.
type DummyResolver struct{}

// Resolve implements resolution.
func (r DummyResolver) Resolve(ctx context.Context, name string) (context.Context, net.IP, error) {
	ctx = saveOriginalAddress(ctx, name)
	return ctx, net.IP{}, nil
}
