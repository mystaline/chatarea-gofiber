package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/mystaline/chatarea-gofiber/internal/app/dto"
	"github.com/mystaline/chatarea-gofiber/internal/app/provider"
	"github.com/mystaline/chatarea-gofiber/internal/app/service"
	"github.com/mystaline/chatarea-gofiber/internal/app/usecase"
	"github.com/mystaline/chatarea-gofiber/internal/app/utils"
)

func GetMyRooms(c *fiber.Ctx) error {
	var response []dto.MyRoom
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

	if query.IsManaged == "true" {
		filterQuery["rooms.creator_id"] = utils.GetExactMatchFilter(user.ID)
	}

	if _, err := useCase.Invoke(usecase.GetMyRoomsParams{
		Context:  c,
		Response: &response,
		Options: service.ServiceOption{
			Filter: filterQuery,
			Joins:  []string{"JOIN rooms ON rooms.id = room_members.room_id", "JOIN users ON users.id = rooms.creator_id"},
			Select: utils.ExtractSelectColumns[dto.MyRoom](),
		},
	}); err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).
			JSON(fiber.Map{"message": err, "status": fiber.ErrInternalServerError.Code})
	}

	fmt.Println("response", response)

	var data []dto.GetMyRoomsResponse
	for _, each := range response {
		data = append(data, dto.GetMyRoomsResponse{
			ID:      each.ID,
			RoomID:  each.RoomID,
			Name:    each.Name,
			Type:    each.Type,
			Address: each.Address,
			Creator: dto.Creator{
				ID:   each.CreatorID,
				Name: each.Name,
			},
			CreatedAt: each.CreatedAt,
			UpdatedAt: each.UpdatedAt,
		})
	}

	return c.JSON(fiber.Map{
		"message": "Successfully get my rooms",
		"data":    data,
	})
}

func JoinRoom(c *fiber.Ctx) error {
	panic("Not implemented")
}

func LeaveRoom(c *fiber.Ctx) error {
	panic("Not implemented")
}
