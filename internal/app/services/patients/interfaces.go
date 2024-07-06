package patients

import (
	"context"
	"konsulin-service/internal/app/models"
	"konsulin-service/internal/pkg/dto/requests"
)

type PatientUsecase interface{}

type PatientRepository interface{}

type PatientFhirClient interface {
	CreatePatient(ctx context.Context, patient *requests.PatientFhir) (*models.Patient, error)
	GetPatientByID(ctx context.Context, patientID string) (*models.Patient, error)
}
