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

type GetProfileParams struct {
	Context  *fiber.Ctx
	Options  service.ServiceOption
	Response *dto.GetMyProfileResponse
}

type GetProfileUseCase struct {
	UserService     service.BaseService
	ServiceProvider provider.ServiceProvider
}

func MakeGetProfileUseCase(serviceProvider provider.ServiceProvider) *GetProfileUseCase {
	return &GetProfileUseCase{
		ServiceProvider: serviceProvider,
	}
}

func (u *GetProfileUseCase) InitServices() {
	u.UserService = u.ServiceProvider.MakeService(config.GetDB(), "users")
}

func (u *GetProfileUseCase) Invoke(params GetProfileParams) (bool, error) {
	if err := u.UserService.FindOne(params.Response, params.Context, params.Options); err != nil {
		fmt.Println("❌ User not found")
		return false, errors.New("user not found or has been deleted")
	}

	return true, nil
}
