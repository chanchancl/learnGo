package main

import (
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/rs/zerolog"
)

func main() {

	app := fiber.New(fiber.Config{
		EnablePrintRoutes: true,
	})

	app.Use(etag.New())

	app.Use(func(c *fiber.Ctx) error {
		log.Println("Run A")
		return c.Next()
	})

	// app.Use(func(c *fiber.Ctx) error {
	// 	log.Println("Run B")

	// 	// Stop here, will not Run C
	// 	return nil
	// })

	app.Use(func(c *fiber.Ctx) error {
		log.Println("Run C")

		return c.Next()
	})

	app.Get("/", func(c *fiber.Ctx) error {
		logger := zerolog.Ctx(c.UserContext())

		logger.Info().Msg("Run DDDD")
		log.Println("Run D")

		return c.SendString(spew.Sdump(c))
	})

	log.Fatal(app.Listen(":8001"))
}
