package sshKraken // import "github.com/carlos-py/sshKraken"

import (
	"net"

	"golang.org/x/net/context"
)

type EmptyResolver struct{}

func (EmptyResolver) Resolve(ctx context.Context, name string) (context.Context, net.IP, error) {
	return ctx, nil, nil
}
