package routers

import (
	"konsulin-service/internal/app/services/auth"

	"github.com/gofiber/fiber/v2"
)

func attachAuthRoutes(router fiber.Router, authController *auth.AuthController) {
	router.Post("/register", authController.RegisterPatient)
	router.Post("/login", authController.Login)
}
