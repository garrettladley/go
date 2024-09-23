package main

import (
	"context"
	"fmt"
)

// use an unexported type as the key
//
// why?
// - to prevent collisions with other packages that use context
// - size
//   - an empty struct has a minimum size of zero bytes but may
//     have a size >0 due to padding. the size of an empty struct
//     in Go is implementation-dependent, and it is usually 1 byte
//     or larger to ensure that each instance of a struct has a
//     unique memory address.
type contextKey struct{}

func SetRequestID(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, contextKey{}, requestID)
}

func GetRequestID(ctx context.Context) (requestID string, ok bool) {
	requestID, ok = ctx.Value(contextKey{}).(string)
	return
}

func main() {
	ctx := context.Background()

	requestID, ok := GetRequestID(ctx)
	fmt.Printf("before setting: requestID: '%s' ok: %t\n", requestID, ok)

	v := "request_id"
	ctx = SetRequestID(ctx, v)

	requestID, ok = GetRequestID(ctx)
	fmt.Printf("after setting: requestID: '%s' ok: %t\n", requestID, ok)
	fmt.Printf("same: %t\n", v == requestID)
}
