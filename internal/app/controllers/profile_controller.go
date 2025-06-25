package controllers

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/mystaline/chatarea-gofiber/internal/app/dto"
	"github.com/mystaline/chatarea-gofiber/internal/app/provider"
	"github.com/mystaline/chatarea-gofiber/internal/app/service"
	"github.com/mystaline/chatarea-gofiber/internal/app/usecase"
	"github.com/mystaline/chatarea-gofiber/internal/app/utils"
)

func Me(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*dto.GetMyProfileResponse)
	if !ok {
		return c.Status(fiber.ErrInternalServerError.Code).
			JSON(fiber.Map{"message": "Invalid user context", "status": fiber.ErrInternalServerError.Code})
	}

	return c.JSON(fiber.Map{
		"message": "Get my profile info",
		"data":    user,
	})
}

func EditProfile(c *fiber.Ctx) error {
	var payload dto.EditProfileBody

	user, ok := c.Locals("user").(*dto.GetMyProfileResponse)
	if !ok {
		return c.Status(fiber.ErrInternalServerError.Code).
			JSON(fiber.Map{"message": "Invalid user context", "status": fiber.ErrInternalServerError.Code})
	}

	useCase := usecase.MakeEditProfileUseCase(&provider.MainServiceProvider{})
	useCase.InitServices()

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"message": fiber.ErrBadRequest.Message,
			"status":  fiber.ErrBadRequest.Code,
		})
	}

	if _, err := useCase.Invoke(usecase.EditProfileParams{
		Context:  c,
		Body:     payload,
		Response: user,
	}); err != nil {
		if strings.Contains(err.Error(), "retrieve data") {
			return c.Status(fiber.ErrNotFound.Code).
				JSON(fiber.Map{"message": err.Error(), "status": fiber.ErrNotFound.Code})
		}
		return c.Status(fiber.ErrInternalServerError.Code).
			JSON(fiber.Map{"message": err.Error(), "status": fiber.ErrInternalServerError.Code})
	}

	return c.JSON(fiber.Map{
		"message": "Successfully edit my profile info",
		"data":    user,
	})
}

func DeleteAccount(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*dto.GetMyProfileResponse)
	if !ok {
		return c.Status(fiber.ErrInternalServerError.Code).
			JSON(fiber.Map{"message": "Invalid user context", "status": fiber.ErrInternalServerError.Code})
	}

	useCase := usecase.MakeDeleteAccountUseCase(&provider.MainServiceProvider{})
	useCase.InitServices()

	filterQuery := map[string]utils.EloquentQuery{
		"id": utils.GetExactMatchFilter(user.ID),
	}

	if _, err := useCase.Invoke(usecase.DeleteAccountParams{
		Context: c,
		Options: service.ServiceOption{Filter: filterQuery},
	}); err != nil {
		if strings.Contains(err.Error(), "retrieve data") {
			return c.Status(fiber.ErrNotFound.Code).
				JSON(fiber.Map{"message": err.Error(), "status": fiber.ErrNotFound.Code})
		}
		return c.Status(fiber.ErrInternalServerError.Code).
			JSON(fiber.Map{"message": err.Error(), "status": fiber.ErrInternalServerError.Code})
	}

	return c.JSON(fiber.Map{
		"message": "Successfully delete an account",
	})
}
