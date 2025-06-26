package usecase

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/mystaline/chatarea-gofiber/internal/app/dto"
	"github.com/mystaline/chatarea-gofiber/internal/app/provider"
	"github.com/mystaline/chatarea-gofiber/internal/app/service"
	"github.com/mystaline/chatarea-gofiber/internal/app/utils"
	"github.com/mystaline/chatarea-gofiber/internal/config"
)

type EditProfileParams struct {
	Context  *fiber.Ctx
	Body     dto.EditProfileBody
	Response *dto.GetMyProfileResponse
}

type EditProfileUseCase struct {
	UserService     service.BaseService
	ServiceProvider provider.ServiceProvider
}

func MakeEditProfileUseCase(serviceProvider provider.ServiceProvider) *EditProfileUseCase {
	return &EditProfileUseCase{
		ServiceProvider: serviceProvider,
	}
}

func (u *EditProfileUseCase) InitServices() {
	u.UserService = u.ServiceProvider.MakeService(config.GetDB(), "users")
}

func (u *EditProfileUseCase) Invoke(params EditProfileParams) (bool, error) {
	hashedPassword, err := utils.ValidateNewPassword(params.Response.Password, params.Context, params.Body.Password)

	if err == nil {
		params.Body.Password = hashedPassword
	}

	if err := u.UserService.UpdateOne(params.Response, params.Context, params.Body, service.ServiceOption{
		Select: utils.ExtractSelectColumns[dto.EditProfileBody](),
	}); err != nil {
		return false, errors.New("failed to edit profile")
	}

	return true, nil
}
