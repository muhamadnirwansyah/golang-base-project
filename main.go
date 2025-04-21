package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhamadnirwansyah/authentication-service/internal/api"
	"github.com/muhamadnirwansyah/authentication-service/internal/config"
	"github.com/muhamadnirwansyah/authentication-service/internal/connection"
	"github.com/muhamadnirwansyah/authentication-service/internal/middleware"
	"github.com/muhamadnirwansyah/authentication-service/internal/repository"
	"github.com/muhamadnirwansyah/authentication-service/internal/service"
)

func main() {

	globalConfiguration := config.Get()

	dbConnection := connection.GetDatabase(globalConfiguration.Database)

	accountRepository := repository.NewAccount(dbConnection)

	authService := service.NewAuthentication(globalConfiguration, accountRepository)
	signUpService := service.NewSignUp(globalConfiguration, accountRepository)
	authHandler := middleware.Authenticate(authService)

	app := fiber.New()

	api.NewAuth(app, authHandler, authService)
	api.NewSignUp(app, authHandler, signUpService)

	_ = app.Listen(globalConfiguration.Server.Host + ":" + globalConfiguration.Server.Port)
}
