package controller

import (
	"strconv"

	"github.com/aiteung/musik"
	dhs "github.com/daniferdinandall/be_dhs2"
	"github.com/gofiber/fiber/v2"
)

func Homepage(c *fiber.Ctx) error {
	ipaddr := musik.GetIPaddress()
	return c.JSON(ipaddr)
}

func GetDhs(c *fiber.Ctx) error {
	queryValue := c.Query("npm")
	i, err := strconv.Atoi(queryValue)
	if err != nil {
		// ... handle error
		// panic(err)
		return c.SendString("format params kurang tepat")
	}
	ps := dhs.GetDhsFromNPM(int(i))

	return c.JSON(ps)
}

func GetAllDhs(c *fiber.Ctx) error {

	ps := dhs.GetDhsAll()
	return c.JSON(ps)
}
