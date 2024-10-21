package main

import (
	"context"
)

type callChainKey struct{}

func AppendCallChain(ctx context.Context, value string) context.Context {
	return context.WithValue(ctx, callChainKey{}, append(GetCallChain(ctx), value))
}

func GetCallChain(ctx context.Context) []string {
	return ctx.Value(callChainKey{}).([]string)
}
