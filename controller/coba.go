package controller

import (
	"github.com/aiteung/musik"
	cek "github.com/aiteung/presensi"
	"github.com/daniferdinandall/Pemrograman-3/ws-dani/config"
	"github.com/gofiber/fiber/v2"
)

func Homepage(c *fiber.Ctx) error {
	ipaddr := musik.GetIPaddress()
	return c.JSON(ipaddr)
}

func GetPresensi(c *fiber.Ctx) error {
	ps := cek.GetPresensiCurrentMonth(config.Ulbimongoconn)
	return c.JSON(ps)
}
