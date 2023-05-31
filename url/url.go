package url

import (
	"github.com/daniferdinandall/Pemrograman-3/ws-dani/controller"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/gofiber/websocket/v2"
)

func Web(page *fiber.App) {
	page.Get("/", controller.Homepage)
	page.Get("/docs/*", swagger.HandlerDefault)
	page.Get("/presensidev", controller.GetPresensi)
	page.Get("/dhs", controller.GetDhs)
	page.Get("/dhs-all", controller.GetAllDhs)
	page.Get("/mahasiswa", controller.GetMhs)
	page.Get("/mahasiswa-all", controller.GetAllMhs)
	page.Get("/dosen", controller.GetDosen)
	page.Get("/dosen-all", controller.GetAllDosen)
	page.Get("/matkul", controller.GetMatkul)
	page.Get("/matkul-all", controller.GetAllMatkul)
	page.Post("/api/whatsauth/request", controller.PostWhatsAuthRequest)  //API from user whatsapp message from iteung gowa
	page.Get("/ws/whatsauth/qr", websocket.New(controller.WsWhatsAuthQR)) //websocket whatsauth
	page.Get("/presensi", controller.GetAllPresensi)                      //menampilkan seluruh data presensi
	page.Get("/presensi/:id", controller.GetPresensiID)                   //menampilkan data presensi berdasarkan id
	page.Post("/ins", controller.InsertData)
	page.Put("/upd/:id", controller.UpdateData)               //update data
	page.Delete("/delete/:id", controller.DeletePresensiByID) //delete data
}
