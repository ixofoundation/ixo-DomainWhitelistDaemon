package whitelist_domain

import (
	"ixowhitelistdaemon/database"

	"github.com/gofiber/fiber/v2"
)

func GetAllWhitelisteDomains(c *fiber.Ctx) error {
	result, err := database.GetAllWhitelisteDomains()
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
			"data":    nil,
		})
	}

	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "DomainList",
		"data":    result,
	})
}

func CreateWhitelistedDomain(c *fiber.Ctx) error {
	NewDomain := new(database.WhitelistDomain)

	err := c.BodyParser(NewDomain)
	if err != nil {
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
			"data":    nil,
		})
		return err
	}

	result, err := database.CreateWhitelistedDomain(NewDomain.Name, NewDomain.Url)
	if err != nil {
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
			"data":    nil,
		})
		return err
	}

	c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "",
		"data":    result,
	})
	return nil
}
