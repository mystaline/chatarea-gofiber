package usecase

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/mystaline/chatarea-gofiber/internal/app/dto"
	"github.com/mystaline/chatarea-gofiber/internal/app/provider"
	"github.com/mystaline/chatarea-gofiber/internal/app/service"
	"github.com/mystaline/chatarea-gofiber/internal/config"
)

type GetUsersParams struct {
	Context  *fiber.Ctx
	Options  service.ServiceOption
	Response *[]dto.GetUserListResponse
}

type GetUsersUseCase struct {
	UserService     service.BaseService
	ServiceProvider provider.ServiceProvider
}

func MakeGetUsersUseCase(serviceProvider provider.ServiceProvider) *GetUsersUseCase {
	return &GetUsersUseCase{
		ServiceProvider: serviceProvider,
	}
}

func (u *GetUsersUseCase) InitServices() {
	u.UserService = u.ServiceProvider.MakeService(config.GetDB(), "users")
}

func (u *GetUsersUseCase) Invoke(params GetUsersParams) (bool, error) {
	if err := u.UserService.FindMany(params.Response, params.Context, params.Options); err != nil {
		fmt.Println("❌ no data")
		return false, errors.New("no data")
	}

	return true, nil
}
