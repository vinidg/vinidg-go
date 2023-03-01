package main

import (
	"github.com/charmbracelet/log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	db, err := NewDatabase()
	if err != nil {
		log.Error("Failed to connect to redis: %s", err.Error())
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Get("/:key/:value", func(c *fiber.Ctx) error {
		list := db.SetValue(c.Params("key"), c.Params("value"))
		return c.SendString(list)
	})

	app.Get("/:key", func(c *fiber.Ctx) error {
		val := db.GetValue(c.Params("key"))
		return c.SendString(val)
	})

	app.Listen(":3000")
}
