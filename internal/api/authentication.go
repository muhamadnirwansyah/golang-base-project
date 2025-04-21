package api

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/muhamadnirwansyah/authentication-service/domain"
	"github.com/muhamadnirwansyah/authentication-service/dto"
)

type authenticationApi struct {
	authService domain.AuthenticationService
}

func NewAuth(app *fiber.App, authHandler fiber.Handler, authService domain.AuthenticationService) {

	authApi := authenticationApi{
		authService: authService,
	}

	app.Post("/v1/authenticate", authApi.authenticate)
	app.Post("/v1/authenticate/validate", authHandler, authApi.authenticateValidate)
}

func (a authenticationApi) authenticate(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.AuthenticationRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(http.StatusUnprocessableEntity)
	}
	res, err := a.authService.Authentication(c, req)

	if err != nil {
		if errors.Is(err, domain.ErrorInvalidCredential) {
			return ctx.Status(http.StatusUnauthorized).JSON(dto.NewResponseMessage("Invalid credentials, Please check your username or password !"))
		}
		return ctx.Status(http.StatusInternalServerError).JSON(dto.NewResponseMessage("Internal Server Error has occured."))
	}

	return ctx.Status(http.StatusOK).JSON(dto.NewResponseData[dto.AuthenticationResponse](res))
}

func (a authenticationApi) authenticateValidate(ctx *fiber.Ctx) error {
	userLocal := ctx.Locals("x-account")
	if userLocal == nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.NewResponseMessage("Sorry, the token is not valid !"))
	}
	return ctx.Status(http.StatusOK).JSON(dto.NewResponseData[dto.AccountData](userLocal.(dto.AccountData)))
}
