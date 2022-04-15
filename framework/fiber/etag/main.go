package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/etag"
)

func main() {
	app := fiber.New(fiber.Config{
		EnablePrintRoutes: true,
	})

	app.Use(etag.New())

	app.Use("/", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	log.Fatal(app.Listen(":8001"))
}
