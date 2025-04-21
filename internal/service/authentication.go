package service

import (
	"context"
	"log/slog"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/muhamadnirwansyah/authentication-service/domain"
	"github.com/muhamadnirwansyah/authentication-service/dto"
	"github.com/muhamadnirwansyah/authentication-service/internal/config"
	"golang.org/x/crypto/bcrypt"
)

type authenticationService struct {
	globalConfig      *config.Config
	accountRepository domain.AccountRepository
}

func (a *authenticationService) Authentication(ctx context.Context, req dto.AuthenticationRequest) (dto.AuthenticationResponse, error) {
	account, err := a.accountRepository.FindByEmail(ctx, req.Email)
	if err != nil {
		slog.ErrorContext(ctx, err.Error())
		return dto.AuthenticationResponse{}, err
	}

	if account.ID == 0 {
		return dto.AuthenticationResponse{}, domain.ErrorInvalidCredential
	}

	if err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(req.Password)); err != nil {
		return dto.AuthenticationResponse{}, domain.ErrorInvalidCredential
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":       account.ID,
			"fullname": account.FullName,
			"email":    account.Email,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})
	tokenString, err := token.SignedString([]byte(a.globalConfig.Secret.Jwt))
	if err != nil {
		slog.ErrorContext(ctx, err.Error())
		return dto.AuthenticationResponse{}, err
	}

	return dto.AuthenticationResponse{
		AccessToken: tokenString,
	}, nil
}

func (a *authenticationService) Validate(ctx context.Context, tokenString string) (dto.AccountData, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(a.globalConfig.Secret.Jwt), nil
	})
	if err != nil {
		slog.ErrorContext(ctx, err.Error())
		return dto.AccountData{}, err
	}
	if token.Valid {
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			idFloat, ok := claims["id"].(float64)
			if !ok {
				return dto.AccountData{}, domain.ErrorInvalidCredential
			}
			return dto.AccountData{
				Id:       int64(idFloat),
				FullName: claims["fullname"].(string),
				Email:    claims["email"].(string),
			}, nil
		}
	}
	return dto.AccountData{}, domain.ErrorInvalidCredential
}

func NewAuthentication(glbc *config.Config, accountRepository domain.AccountRepository) domain.AuthenticationService {
	return &authenticationService{
		globalConfig:      glbc,
		accountRepository: accountRepository,
	}
}
