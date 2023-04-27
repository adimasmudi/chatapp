package controllers

import (
	"chatapp/helper"
	"chatapp/inputs"
	"chatapp/services"
	"context"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

type userHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) Register(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var input inputs.RegisterUserInput

	// file, err := c.FormFile("picture")

	// if err != nil{
	// 	response := helper.APIResponse("Wrong file format", http.StatusBadRequest, "error", &fiber.Map{"error" : err})
	// 	c.Status(http.StatusBadRequest).JSON(response)
	// 	return nil
	// }

	//validate the request body
	if err := c.BodyParser(&input); err != nil {
		response := helper.APIResponse("Register Failed", http.StatusBadRequest, "error", &fiber.Map{"error" : err})
		c.Status(http.StatusBadRequest).JSON(response)
		return nil
	}

	// fileToSave := fmt.Sprintf("./images/%s-%s",input.Username,file.Filename)

	registeredUser, err := h.userService.Register(ctx, input)

	if err != nil{
		response := helper.APIResponse("Register Failed", http.StatusBadRequest, "error", &fiber.Map{"error" : err})
		c.Status(http.StatusBadRequest).JSON(response)
		return nil
	}

	// c.SaveFile(file, fileToSave)

	response := helper.APIResponse("User register success", http.StatusOK, "success", registeredUser)
	c.Status(http.StatusOK).JSON(response)
	return nil
}

func (h *userHandler) Login(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var input inputs.LoginUserInput

	//validate the request body
	if err := c.BodyParser(&input); err != nil {
		response := helper.APIResponse("Login Failed ini", http.StatusBadRequest, "error", &fiber.Map{"error" : err})
		c.Status(http.StatusBadRequest).JSON(response)
		return nil
	}

	logedinUser, token,  err := h.userService.Login(ctx,input)

	if err != nil{
		response := helper.APIResponse("Login Failed di", http.StatusBadRequest, "error", &fiber.Map{"error" : err})
		c.Status(http.StatusBadRequest).JSON(response)
		return nil
	}

	// Create cookie
	cookie := new(fiber.Cookie)
	cookie.Name = "user"
	cookie.Value = logedinUser.Username
	cookie.Expires = time.Now().Add(24 * time.Hour)

	// Set cookie
	c.Cookie(cookie)

	response := helper.APIResponse("Login success", http.StatusOK, "success", &fiber.Map{"user" : logedinUser, "token" : token})
	c.Status(http.StatusOK).JSON(response)
	return nil
}