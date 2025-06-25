package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/mystaline/chatarea-gofiber/internal/app/dto"
	"github.com/mystaline/chatarea-gofiber/internal/app/models"
	"github.com/mystaline/chatarea-gofiber/internal/app/provider"
	"github.com/mystaline/chatarea-gofiber/internal/app/service"
	"github.com/mystaline/chatarea-gofiber/internal/app/usecase"
	"github.com/mystaline/chatarea-gofiber/internal/app/utils"
)

func GetMyRooms(c *fiber.Ctx) error {
	var response []models.UserRoom
	user, ok := c.Locals("user").(*dto.GetMyProfileResponse)
	if !ok {
		return c.Status(fiber.ErrInternalServerError.Code).
			JSON(fiber.Map{"message": "Invalid user context", "status": fiber.ErrInternalServerError.Code})
	}

	useCase := usecase.MakeGetMyRoomsUseCase(&provider.MainServiceProvider{})
	useCase.InitServices()

	var query dto.RoomQuery
	if err := c.QueryParser(&query); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid query",
			"status":  "400",
		})
	}

	filterQuery := map[string]utils.EloquentQuery{
		"user_id": utils.GetExactMatchFilter(user.ID),
	}

	if query.IsManaged != "" {
		filterQuery["room.creator_id"] = utils.GetExactMatchFilter(user.ID)
	}

	if _, err := useCase.Invoke(usecase.GetMyRoomsParams{
		Context:  c,
		Response: &response,
		Options: service.ServiceOption{
			Filter:  filterQuery,
			Preload: []string{"User", "Room", "Room.Creator"},
		},
	}); err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).
			JSON(fiber.Map{"message": err, "status": fiber.ErrInternalServerError.Code})
	}

	fmt.Println("response", response)

	return c.JSON(fiber.Map{
		"message": "Successfully get my rooms",
		"data":    response,
	})
}

func JoinRoom(c *fiber.Ctx) error {
	panic("Not implemented")
}

func LeaveRoom(c *fiber.Ctx) error {
	panic("Not implemented")
}
