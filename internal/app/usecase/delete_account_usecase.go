package usecase

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/mystaline/chatarea-gofiber/internal/app/provider"
	"github.com/mystaline/chatarea-gofiber/internal/app/service"
	"github.com/mystaline/chatarea-gofiber/internal/config"
)

type DeleteAccountParams struct {
	Context *fiber.Ctx
	Options service.ServiceOption
}

type DeleteAccountUseCase struct {
	UserService     service.BaseService
	ServiceProvider provider.ServiceProvider
}

func MakeDeleteAccountUseCase(serviceProvider provider.ServiceProvider) *DeleteAccountUseCase {
	return &DeleteAccountUseCase{
		ServiceProvider: serviceProvider,
	}
}

func (u *DeleteAccountUseCase) InitServices() {
	u.UserService = u.ServiceProvider.MakeService(config.GetDB(), "users")
}

func (u *DeleteAccountUseCase) Invoke(params DeleteAccountParams) (bool, error) {
	if err := u.UserService.DeleteOne(params.Context, params.Options); err != nil {
		fmt.Println("❌ User not found")
		return false, errors.New("user not found or has been deleted")
	}

	return true, nil
}
