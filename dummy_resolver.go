package dphx

import (
	"net"

	"golang.org/x/net/context"
)

// DummyResolver stores unresolved host address at context for later use.
type DummyResolver struct{}

// OriginalAddress preserves raw address for delayed DNS resolution.
type OriginalAddress string

func (o OriginalAddress) String() string {
	return string(o)
}

// Resolve implements resolution.
func (r DummyResolver) Resolve(ctx context.Context, name string) (context.Context, net.IP, error) {
	ctx = context.WithValue(ctx, "originalAddress", OriginalAddress(name))
	return ctx, net.IP{}, nil
}
