package auth

import (
	"context"
	"konsulin-service/internal/pkg/dto/requests"
	"konsulin-service/internal/pkg/dto/responses"
)

type AuthUsecase interface {
	RegisterPatient(ctx context.Context, request *requests.RegisterPatient) (*responses.RegisterPatient, error)
	LoginPatient(ctx context.Context, request *requests.LoginPatient) (*responses.LoginPatient, error)
}

type AuthRepository interface{}
