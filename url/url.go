package url

import (
	"github.com/daniferdinandall/Pemrograman-3/ws-dani/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func Web(page *fiber.App) {
	page.Get("/", controller.Homepage)
	page.Get("/dhs", controller.GetDhs)
	page.Get("/dhs-all", controller.GetAllDhs)
	page.Get("/presensi", controller.GetPresensi)
	page.Post("/api/whatsauth/request", controller.PostWhatsAuthRequest)  //API from user whatsapp message from iteung gowa
	page.Get("/ws/whatsauth/qr", websocket.New(controller.WsWhatsAuthQR)) //websocket whatsauth

}
