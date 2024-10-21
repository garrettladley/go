package main

import (
	"context"
	"fmt"
)

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
