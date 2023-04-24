package main

import (
	"chatapp/configs"
	"chatapp/ws"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())
	app.Use(logger.New())
	app.Use(recover.New())

	configs.ConnectDB()

	hub := ws.NewHub()
	wsHandler := ws.NewHandler(hub)

	go hub.Run()

	app.Post("/ws/createRoom", wsHandler.CreateRoom)
	app.Get("/ws/joinRoom/:roomId", ws.JoinRoom(hub))
	app.Get("/ws/getRooms",wsHandler.GetRooms)
	app.Get("/ws/getClients/:roomId", wsHandler.GetClients)

	app.Listen(":5000")
}