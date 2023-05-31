package main

import (
	"log"
	"os"

	"github.com/daniferdinandall/Pemrograman-3/ws-dani/config"

	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/whatsauth/whatsauth"

	"github.com/daniferdinandall/Pemrograman-3/ws-dani/url"

	"github.com/gofiber/fiber/v2"
)

func Dangdut() (port string) {
	port = os.Getenv("PORT")
	if port == "" {
		port = ":8000"
	} else if port[0:1] != ":" {
		port = ":" + port
	}
	return
}

// @title TES SWAG
// @version 1.0
// @description This is a sample swagger for Fiber

// @contact.name API Support
// @contact.url https://github.com/indrariksa
// @contact.email indra@ulbi.ac.id

// @host ws-ulbi.herokuapp.com
// @BasePath /
// @schemes https http

func main() {
	go whatsauth.RunHub()
	site := fiber.New(config.Iteung)
	site.Use(cors.New(config.Cors))
	url.Web(site)
	log.Fatal(site.Listen(Dangdut()))
}
