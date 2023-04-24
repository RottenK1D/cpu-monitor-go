package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", root)
	app.Static("/static", "static")

	log.Fatal(app.Listen(":3000"))
}

func root(c *fiber.Ctx) error {
	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendFile("templates/index.html")
}
