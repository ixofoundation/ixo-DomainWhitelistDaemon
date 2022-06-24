package main

import (
	"ixowhitelistdaemon/database"
	whitelist_domain "ixowhitelistdaemon/whitelist"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func status(c *fiber.Ctx) error {
	return c.SendString("Server is running! Send your request")
}

func setupRoutes(app *fiber.App) {

	app.Get("/", status)

	app.Get("/api/getwhitelist", whitelist_domain.GetAllWhitelisteDomains)
	app.Post("/api/createwhitelistitem", whitelist_domain.CreateWhitelistedDomain)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	app := fiber.New()
	dbErr := database.InitDatabase()

	if dbErr != nil {
		panic(dbErr)
	}

	setupRoutes(app)
	app.Listen(":3000")
}
