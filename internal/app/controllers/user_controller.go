package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mystaline/chatarea-gofiber/internal/app/dto"
	"github.com/mystaline/chatarea-gofiber/internal/app/provider"
	"github.com/mystaline/chatarea-gofiber/internal/app/service"
	"github.com/mystaline/chatarea-gofiber/internal/app/usecase"
	"github.com/mystaline/chatarea-gofiber/internal/app/utils"
)

func GetUsers(c *fiber.Ctx) error {
	var users []dto.GetUserListResponse

	useCase := usecase.MakeGetUsersUseCase(&provider.MainServiceProvider{})
	useCase.InitServices()

	var query dto.UserQuery
	if err := c.QueryParser(&query); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid query",
			"status":  "400",
		})
	}

	filterQuery := map[string]utils.EloquentQuery{
		"id":       utils.GetExactMatchFilter(query.ID),
		"name":     utils.GetExactMatchFilter(query.Name),
		"username": utils.GetExactMatchFilter(query.Username),
	}

	_, err := useCase.Invoke(usecase.GetUsersParams{
		Context:  c,
		Response: &users,
		Options:  service.ServiceOption{Filter: filterQuery},
	})

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(users)
}
