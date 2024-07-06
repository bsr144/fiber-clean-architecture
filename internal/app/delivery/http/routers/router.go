package routers

import (
	"fmt"
	"konsulin-service/internal/app/config"
	"konsulin-service/internal/app/services/auth"
	"konsulin-service/internal/app/services/patients"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(
	app *fiber.App,
	appConfig config.App,
	patientController *patients.PatientController,
	authController *auth.AuthController,
) {
	endpointPrefix := fmt.Sprintf("/%s", appConfig.EndpointPrefix)
	api := app.Group(endpointPrefix)

	versionPrefix := fmt.Sprintf("/%s", appConfig.Version)
	apiVersion := api.Group(versionPrefix)

	authRouter := apiVersion.Group("/auth")
	patientRouter := apiVersion.Group("/patients")

	attachPatientRoutes(patientRouter, patientController)
	attachAuthRoutes(authRouter, authController)
}
