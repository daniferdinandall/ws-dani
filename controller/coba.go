package controller

import (
	"strconv"

	"github.com/aiteung/musik"
	cek "github.com/aiteung/presensi"
	"github.com/daniferdinandall/Pemrograman-3/ws-dani/config"
	"github.com/gofiber/fiber/v2"

	dhs "github.com/daniferdinandall/be_dhs2"
)

func Homepage(c *fiber.Ctx) error {
	ipaddr := musik.GetIPaddress()
	return c.JSON(ipaddr)
}

func GetPresensi(c *fiber.Ctx) error {
	ps := cek.GetPresensiCurrentMonth(config.Ulbimongoconn)
	return c.JSON(ps)
}

// dhs

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

//mahasiswa

func GetMhs(c *fiber.Ctx) error {
	queryValue := c.Query("npm")
	i, err := strconv.Atoi(queryValue)
	if err != nil {
		// ... handle error
		// panic(err)
		return c.SendString("format params kurang tepat")
	}
	ps := dhs.GetMhsFromNPM(int(i))

	return c.JSON(ps)
}

func GetAllMhs(c *fiber.Ctx) error {

	ps := dhs.GetMhsAll()
	return c.JSON(ps)
}

// Dosen

func GetDosen(c *fiber.Ctx) error {
	queryValue := c.Query("kode")

	ps := dhs.GetDosenFromKodeDosen(queryValue)

	return c.JSON(ps)
}

func GetAllDosen(c *fiber.Ctx) error {

	ps := dhs.GetDosenAll()
	return c.JSON(ps)
}

func GetMatkul(c *fiber.Ctx) error {
	queryValue := c.Query("kode")

	ps := dhs.GetMatkulFromKodeMatkul(queryValue)

	return c.JSON(ps)
}

func GetAllMatkul(c *fiber.Ctx) error {

	ps := dhs.GetMatkulAll()
	return c.JSON(ps)
}
