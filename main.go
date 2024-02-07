package main

import (
	pub "gcp-pub-sub/publisher"
	sub "gcp-pub-sub/subscriber"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	go func() {
		sub.Sub()
	}()

	app.Post("/", func(c *fiber.Ctx) error {
		pub.Pub()
		return c.SendString("publish message successfully")
	})

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
