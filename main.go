package main

import (
	"fmt"

	"github.com/gofiber/fiber"
)

func main() {
	fmt.Print("hello world")
	app := fiber.New()
	// get API
	// app.Get("/abc", func(c *fiber.Ctx) error {
	// 	return c.SendString("hello world")
	// })
	app.Get("/abc", func(c *fiber.Ctx) {
	
	c.SendString("first API \n second sentence")
	})
	app.Listen(":3000")
}
