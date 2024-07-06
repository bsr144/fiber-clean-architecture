package utils

import (
	"konsulin-service/internal/pkg/constvars"
	"konsulin-service/internal/pkg/dto/requests"
)

func BuildFhirPatientRequest(username, email string) *requests.PatientFhir {
	return &requests.PatientFhir{
		ResourceType: constvars.ResourcePatient,
		Telecom: []requests.ContactPoint{
			{
				System: "email",
				Value:  email,
				Use:    "home",
			},
		},
	}
}
