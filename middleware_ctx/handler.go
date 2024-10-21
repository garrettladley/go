package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Handler(c *fiber.Ctx) error {
	fmt.Println(GetCallChain(c.Context()))

	// get the client from the context
	client := GetFoo(c)
	client.DoSomething("Handler")

	// call another func
	OtherFunc(c.Context())

	return c.SendString("Hello, World!")
}
