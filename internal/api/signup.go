package api

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/muhamadnirwansyah/authentication-service/domain"
	"github.com/muhamadnirwansyah/authentication-service/dto"
)

type signUpApi struct {
	signupService domain.SignUpService
}

func NewSignUp(app *fiber.App, authHandler fiber.Handler, signupService domain.SignUpService) {

	signUpApi := signUpApi{
		signupService: signupService,
	}

	app.Post("/v1/signup", signUpApi.signUpNewAccount)
	app.Put("/v1/update", authHandler, signUpApi.updateAccount)
}

func (s signUpApi) updateAccount(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.UpdateAccountRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(dto.NewResponseMessage("Please check your payload request !"))
	}

	res, err := s.signupService.UpddateAccount(c, req)
	if err != nil {
		if errors.Is(err, domain.ErrorAccountNotFound) {
			return ctx.Status(http.StatusNotFound).JSON(dto.NewResponseMessage("Account not found"))
		}
		if errors.Is(err, domain.ErrorEmailIsAlreadyExists) {
			return ctx.Status(http.StatusConflict).JSON(dto.NewResponseMessage("Email is already use"))
		}
		log.Fatalf("Update account error : %v", err)
		return ctx.Status(http.StatusInternalServerError).JSON(dto.NewResponseMessage("something went wrong"))
	}

	return ctx.Status(http.StatusOK).JSON(dto.NewResponseData[dto.UpdateAccountResponse](res))
}

func (s signUpApi) signUpNewAccount(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.SignUpRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(dto.NewResponseMessage("Please check your payload request !"))
	}

	res, err := s.signupService.SignUp(c, req)
	if err != nil {
		if errors.Is(err, domain.ErrorEmailIsAlreadyExists) {
			return ctx.Status(http.StatusConflict).JSON(dto.NewResponseMessage("Email is already use !"))
		}
		/* fallback unknown error **/
		return ctx.Status(http.StatusInternalServerError).JSON(dto.NewResponseMessage("Something went wrong !"))
	}

	return ctx.Status(http.StatusOK).JSON(dto.NewResponseData[dto.SignUpResponse](res))
}
