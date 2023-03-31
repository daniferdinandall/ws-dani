package controller

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/aiteung/musik"
	cek "github.com/aiteung/presensi"
	"github.com/daniferdinandall/Pemrograman-3/ws-dani/config"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	module "github.com/daniferdinandall/be_dhs2/module"
	modulePresensi "github.com/indrariksa/be_presensi/module"
)

func Homepage(c *fiber.Ctx) error {
	ipaddr := musik.GetIPaddress()
	return c.JSON(ipaddr)
}
func GetPresensiID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Wrong parameter",
		})
	}
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "Invalid id parameter",
		})
	}
	ps, err := modulePresensi.GetPresensiFromID(objID, config.Ulbimongoconn, "presensi")
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"status":  http.StatusNotFound,
				"message": fmt.Sprintf("No data found for id %s", id),
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Error retrieving data for id %s", id),
		})
	}
	return c.JSON(ps)
}

func GetAllPresensi(c *fiber.Ctx) error {
	ps := modulePresensi.GetAllPresensi(config.Ulbimongoconn, "presensi")
	return c.JSON(ps)
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
	ps := module.GetDhsFromNPM(int(i))

	return c.JSON(ps)
}

func GetAllDhs(c *fiber.Ctx) error {

	ps := module.GetDhsAll()
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
	ps := module.GetMhsFromNPM(int(i))

	return c.JSON(ps)
}

func GetAllMhs(c *fiber.Ctx) error {

	ps := module.GetMhsAll()
	return c.JSON(ps)
}

// Dosen

func GetDosen(c *fiber.Ctx) error {
	queryValue := c.Query("kode")

	ps := module.GetDosenFromKodeDosen(queryValue)

	return c.JSON(ps)
}

func GetAllDosen(c *fiber.Ctx) error {

	ps := module.GetDosenAll()
	return c.JSON(ps)
}

func GetMatkul(c *fiber.Ctx) error {
	queryValue := c.Query("kode")

	ps := module.GetMatkulFromKodeMatkul(queryValue)

	return c.JSON(ps)
}

func GetAllMatkul(c *fiber.Ctx) error {

	ps := module.GetMatkulAll()
	return c.JSON(ps)
}
