package main

import "github.com/gofiber/fiber/v2"

func Middleware0(c *fiber.Ctx) error {
	c.SetUserContext(AppendCallChain(c.Context(), "middleware0"))
	return c.Next()
}

func Middleware1(c *fiber.Ctx) error {
	c.SetUserContext(AppendCallChain(c.Context(), "middleware1"))
	return c.Next()
}

func Middleware2(c *fiber.Ctx) error {
	c.SetUserContext(AppendCallChain(c.Context(), "middleware2"))
	return c.Next()
}
