package service

import (
	"context"
	"log"

	"github.com/muhamadnirwansyah/authentication-service/domain"
	"github.com/muhamadnirwansyah/authentication-service/dto"
	"github.com/muhamadnirwansyah/authentication-service/internal/config"
	"golang.org/x/crypto/bcrypt"
)

func NewSignUp(glbc *config.Config, accountRepository domain.AccountRepository) domain.SignUpService {
	return &signUpService{
		globalConfig:      glbc,
		accountRepository: accountRepository,
	}
}

type signUpService struct {
	globalConfig      *config.Config
	accountRepository domain.AccountRepository
}

func (a *signUpService) UpddateAccount(ctx context.Context, req dto.UpdateAccountRequest) (dto.UpdateAccountResponse, error) {

	existingAccount, err := a.accountRepository.FindById(ctx, req.ID)
	if err == nil && existingAccount.ID == 0 {
		return dto.UpdateAccountResponse{}, domain.ErrorAccountNotFound
	}

	existingEmail, err := a.accountRepository.FindByEmail(ctx, req.Email)
	if err == nil && existingEmail.Email != "" {
		return dto.UpdateAccountResponse{}, domain.ErrorEmailIsAlreadyExists
	}

	existingAccount.ID = req.ID
	existingAccount.Email = req.Email
	existingAccount.FullName = req.FullName
	existingAccount.PhoneNumber = req.PhoneNumber
	hashPassword, err := hashPasswordToBycrypt(req.Password)
	if err != nil {
		return dto.UpdateAccountResponse{}, domain.ErrorInternalServerError
	}
	existingAccount.Password = hashPassword

	err = a.accountRepository.Update(ctx, &existingAccount)
	if err != nil {
		log.Fatalf("error execute repository update : %v", err)
		return dto.UpdateAccountResponse{}, domain.ErrorInternalServerError
	}

	response := dto.UpdateAccountResponse{
		ID:          existingAccount.ID,
		FullName:    existingAccount.FullName,
		PhoneNumber: existingAccount.PhoneNumber,
		Email:       existingAccount.Email,
	}
	return response, nil
}

func hashPasswordToBycrypt(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (a *signUpService) SignUp(ctx context.Context, req dto.SignUpRequest) (dto.SignUpResponse, error) {

	existing, err := a.accountRepository.FindByEmail(ctx, req.Email)
	if err == nil && existing.Email != "" {
		return dto.SignUpResponse{}, domain.ErrorEmailIsAlreadyExists
	}

	hashPassword, err := hashPasswordToBycrypt(req.Password)
	if err != nil {
		return dto.SignUpResponse{}, domain.ErrorInternalServerError
	}

	newAccount := &domain.Account{
		FullName: req.FullName, Email: req.Email, PhoneNumber: req.PhoneNumber, Password: hashPassword,
	}

	if err := a.accountRepository.Save(ctx, newAccount); err != nil {
		return dto.SignUpResponse{}, domain.ErrorInternalServerError
	}

	return dto.SignUpResponse{
		FullName: req.FullName, Email: req.Email, PhoneNumber: req.PhoneNumber,
	}, nil
}
