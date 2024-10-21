package main

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type fooKey struct{}

func SetFoo(ctx context.Context, client *fooClient) context.Context {
	return context.WithValue(ctx, fooKey{}, client)
}

func GetFoo(c *fiber.Ctx) *fooClient {
	return c.Context().Value(fooKey{}).(*fooClient)
}

// Foo is a third-party service we are using
type fooClient struct{}

func (f *fooClient) DoSomething(from string) {
	fmt.Println("fooClient.DoSomething called from", from)
}

func WithFoo(c *fiber.Ctx) error {
	AppendCallChain(c.Context(), "WithFoo")
	// set up code ..

	// configure the client
	client := &fooClient{}

	// store the client in the context
	SetFoo(c.Context(), client)

	// next middleware
	return c.Next()
}
