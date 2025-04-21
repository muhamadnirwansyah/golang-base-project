package domain

import (
	"context"

	"github.com/muhamadnirwansyah/authentication-service/dto"
)

type SignUpService interface {
	SignUp(ctx context.Context, req dto.SignUpRequest) (dto.SignUpResponse, error)
	UpddateAccount(ctx context.Context, req dto.UpdateAccountRequest) (dto.UpdateAccountResponse, error)
}
