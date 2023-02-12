package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func home(c *fiber.Ctx) error {
	return c.SendFile("./content/main.html")
}

func api(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		".name":    "Aryaan",
		"about":    "Backend Developer (Golang, Java, Python) & Data Science Enthusiast",
		"email":    "avsheth03@gmail.com",
		"github":   "https://github.com/AryaanSheth",
		"linkedin": "https://www.linkedin.com/in/aryaan-sheth-b9565423b/",
	})
}

func main() {
	app := fiber.New()

	go app.Get("/api", api)

	go app.Get("/", home)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(app.Listen("0.0.0.0:" + port))

}
