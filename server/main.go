package main

import (
	"chatapp/configs"
	"chatapp/routes"
	"chatapp/ws"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
        AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin,Authorization",
        AllowOrigins:     "http://localhost:3000",
        AllowCredentials: true,
        AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
    }))
	app.Use(logger.New())
	app.Use(recover.New())

	configs.ConnectDB()

	// collections
	var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")


	api := app.Group("/api/v1")

	// routes
	routes.UserRoute(api, userCollection)

	hub := ws.NewHub()
	wsHandler := ws.NewHandler(hub)

	go hub.Run()

	app.Post("/ws/createRoom", wsHandler.CreateRoom)
	app.Get("/ws/joinRoom/:roomId", ws.JoinRoom(hub))
	app.Get("/ws/getRooms",wsHandler.GetRooms)
	app.Get("/ws/getClients/:roomId", wsHandler.GetClients)

	app.Listen(":5000")
}