package main

import (
	"context"
	"fmt"
)

func Example_ctx_get_before_set() {
	ctx := context.Background()

	requestID, ok := GetRequestID(ctx)
	fmt.Printf("before setting: requestID: '%s' ok: %t", requestID, ok)
	// Output: before setting: requestID: '' ok: false
}

func Example_ctx_get_after_set() {
	ctx := context.Background()

	v := "request_id"
	ctx = SetRequestID(ctx, v)

	requestID, ok := GetRequestID(ctx)
	fmt.Printf("after setting: requestID: '%s' ok: %t", requestID, ok)
	// Output: after setting: requestID: 'request_id' ok: true
}

func Example_ctx_set_get_same() {
	ctx := context.Background()

	v := "request_id"
	ctx = SetRequestID(ctx, v)

	requestID, _ := GetRequestID(ctx)
	fmt.Printf("same: %t", v == requestID)
	// Output: same: true
}
