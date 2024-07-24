package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Static("/", "./assets")
	app.Use(logger.New())
	app.Listen(":3005")
}
