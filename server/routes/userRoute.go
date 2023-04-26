package routes

import (
	"chatapp/controllers"
	"chatapp/repositories"
	"chatapp/services"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func UserRoute(api fiber.Router, userCollection *mongo.Collection) {

	userRepository := repositories.NewUserRepository(userCollection)
	userService := services.NewUserService(userRepository)
	userHandler := controllers.NewUserHandler(userService)

	authUser := api.Group("/auth/users")

	authUser.Post("/register", userHandler.Register)
	authUser.Post("/login", userHandler.Login)

}