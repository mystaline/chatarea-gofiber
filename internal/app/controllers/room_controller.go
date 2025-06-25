package controllers

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/mystaline/chatarea-gofiber/internal/app/dto"
	"github.com/mystaline/chatarea-gofiber/internal/app/models"
	"github.com/mystaline/chatarea-gofiber/internal/app/provider"
	"github.com/mystaline/chatarea-gofiber/internal/app/usecase"
)

func CreateRoom(c *fiber.Ctx) error {
	var response dto.CreateRoomResponse
	var body dto.CreateRoomBody

	useCase := usecase.MakeCreateRoomUseCase(&provider.MainServiceProvider{})
	useCase.InitServices()

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"message": fiber.ErrBadRequest.Message,
			"status":  fiber.ErrBadRequest.Code,
		})
	}

	if _, err := useCase.Invoke(usecase.CreateRoomParams{
		Context:  c,
		Body:     body,
		Response: response,
	}); err != nil {
		if strings.Contains(err.Error(), "user not found") {
			return c.Status(fiber.ErrNotFound.Code).
				JSON(fiber.Map{"message": "There's no user with that id", "status": fiber.ErrNotFound.Code})
		}
		return c.Status(fiber.ErrInternalServerError.Code).
			JSON(fiber.Map{"message": "Failed to create room", "error": err.Error(), "status": fiber.ErrInternalServerError.Code})
	}

	return c.JSON(fiber.Map{
		"message": "Successfully create room",
		"data":    response,
	})
}

func GetRoomInfo(c *fiber.Ctx) error {
	var response models.Room

	useCase := usecase.MakeGetRoomInfoUseCase(&provider.MainServiceProvider{})
	useCase.InitServices()

	if _, err := useCase.Invoke(usecase.GetRoomInfoParams{
		Context:  c,
		Response: &response,
	}); err != nil {
		return c.Status(fiber.ErrNotFound.Code).
			JSON(fiber.Map{"message": "Room doesn't exist", "status": fiber.ErrNotFound.Code})
	}

	return c.JSON(fiber.Map{
		"message": "Successfully get room info",
		"data":    response,
	})
}

func EditRoomInfo(c *fiber.Ctx) error {
	var body dto.EditRoomBody

	useCase := usecase.MakeEditRoomUseCase(&provider.MainServiceProvider{})
	useCase.InitServices()

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"message": fiber.ErrBadRequest.Message,
			"status":  fiber.ErrBadRequest.Code,
		})
	}

	if _, err := useCase.Invoke(usecase.EditRoomParams{
		Context: c,
		Body:    body,
	}); err != nil {
		if strings.Contains(err.Error(), "room not found") {
			return c.Status(fiber.ErrNotFound.Code).
				JSON(fiber.Map{"message": "Room doesn't exist", "status": fiber.ErrNotFound.Code})
		}
		return c.Status(fiber.ErrInternalServerError.Code).
			JSON(fiber.Map{"message": "Failed to edit room", "error": err.Error(), "status": fiber.ErrInternalServerError.Code})
	}

	return c.JSON(fiber.Map{
		"message": "Successfully edit room",
	})
}

func DeleteRoom(c *fiber.Ctx) error {
	useCase := usecase.MakeDeleteRoomUseCase(&provider.MainServiceProvider{})
	useCase.InitServices()

	if _, err := useCase.Invoke(usecase.DeleteRoomParams{
		Context: c,
	}); err != nil {
		return c.Status(fiber.ErrNotFound.Code).
			JSON(fiber.Map{"message": "Room doesn't exist", "status": fiber.ErrNotFound.Code})
	}

	return c.JSON(fiber.Map{
		"message": "Successfully get room info",
	})
}
