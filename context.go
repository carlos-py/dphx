package dphx

import (
	"golang.org/x/net/context"
)

// OriginalAddress preserves raw address for delayed DNS resolution.
type OriginalAddress string

func (o OriginalAddress) String() string {
	return string(o)
}

func saveOriginalAddress(ctx context.Context, name string) context.Context {
	return context.WithValue(ctx, "originalAddress", OriginalAddress(name))
}

func fetchOriginalAddress(ctx context.Context) string {
	originalAddress := ctx.Value("originalAddress")
	if originalAddress == nil {
		return ""
	}

	originalAddressTyped, ok := originalAddress.(OriginalAddress)
	if !ok {
		return ""
	}

	return originalAddressTyped.String()
}
