package routers

import (
	"konsulin-service/internal/app/services/patients"

	"github.com/gofiber/fiber/v2"
)

func attachPatientRoutes(router fiber.Router, patientController *patients.PatientController) {
	// router.Post("", patientController.CreatePatient)
}
