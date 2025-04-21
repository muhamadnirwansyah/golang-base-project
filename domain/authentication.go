package domain

import (
	"context"

	"github.com/muhamadnirwansyah/authentication-service/dto"
)

type AuthenticationService interface {
	Authentication(ctx context.Context, req dto.AuthenticationRequest) (dto.AuthenticationResponse, error)
	Validate(ctx context.Context, tokenString string) (dto.AccountData, error)
}
