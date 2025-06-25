package controllers

import (
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/mystaline/chatarea-gofiber/internal/app/dto"
	"github.com/mystaline/chatarea-gofiber/internal/app/provider"
	"github.com/mystaline/chatarea-gofiber/internal/app/service"
	"github.com/mystaline/chatarea-gofiber/internal/app/usecase"
	"github.com/mystaline/chatarea-gofiber/internal/app/utils"
)

func Login(c *fiber.Ctx) error {
	var payload dto.LoginBody
	var response dto.LoginResponse

	useCase := usecase.MakeLoginUserUseCase(&provider.MainServiceProvider{})
	useCase.InitServices()

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"message": fiber.ErrBadRequest.Message,
			"status":  fiber.ErrBadRequest.Code,
		})
	}

	validate := validator.New()
	if err := validate.Struct(&payload); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"error":   err.Error(),
			"message": fiber.ErrBadRequest.Message,
			"status":  fiber.ErrBadRequest.Code,
		})
	}

	filterQuery := map[string]utils.EloquentQuery{
		"username": utils.GetExactMatchFilter(payload.Username),
	}

	_, err := useCase.Invoke(usecase.LoginUserParams{
		Context:  c,
		Body:     payload,
		Response: &response,
		Options:  service.ServiceOption{Filter: filterQuery},
	})

	if err != nil {
		if strings.Contains(err.Error(), "user not found") {
			return c.Status(fiber.ErrNotFound.Code).JSON(fiber.Map{
				"message": err.Error(),
				"status":  fiber.ErrNotFound.Code,
			})
		}

		return c.Status(fiber.ErrUnauthorized.Code).JSON(fiber.Map{
			"message": err.Error(),
			"status":  fiber.ErrUnauthorized.Code,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  fiber.StatusCreated,
		"message": "Successfully logged in",
		"token":   response.Token,
	})
}

func Register(c *fiber.Ctx) error {
	var payload dto.RegisterBody

	useCase := usecase.MakeRegisterUserUseCase(&provider.MainServiceProvider{})
	useCase.InitServices()

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"message": fiber.ErrBadRequest.Message,
			"status":  fiber.ErrBadRequest.Code,
		})
	}

	validate := validator.New()
	if err := validate.Struct(&payload); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"error":   err.Error(),
			"message": fiber.ErrBadRequest.Message,
			"status":  fiber.ErrBadRequest.Code,
		})
	}

	filterQuery := map[string]utils.EloquentQuery{
		"username": utils.GetExactMatchFilter(payload.Username),
	}

	_, err := useCase.Invoke(usecase.RegisterUserParams{
		Context: c,
		Body:    payload,
		Options: service.ServiceOption{Filter: filterQuery},
	})

	if err != nil {
		if strings.Contains(err.Error(), "username already used") {
			return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
				"message": err.Error(),
				"status":  fiber.ErrBadRequest.Code,
			})
		}

		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"message": err.Error(),
			"status":  fiber.ErrInternalServerError.Code,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  fiber.StatusCreated,
		"message": "Successfully register an user",
		"data":    payload,
	})
}
