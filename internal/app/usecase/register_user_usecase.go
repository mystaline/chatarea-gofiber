package usecase

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/mystaline/chatarea-gofiber/internal/app/dto"
	"github.com/mystaline/chatarea-gofiber/internal/app/provider"
	"github.com/mystaline/chatarea-gofiber/internal/app/service"
	"github.com/mystaline/chatarea-gofiber/internal/app/utils"
	"github.com/mystaline/chatarea-gofiber/internal/config"
	"golang.org/x/crypto/bcrypt"
)

type RegisterUserParams struct {
	Context *fiber.Ctx
	Options service.ServiceOption
	Body    dto.RegisterBody
}

type RegisterUserUseCase struct {
	UserService     service.BaseService
	ServiceProvider provider.ServiceProvider
}

func MakeRegisterUserUseCase(serviceProvider provider.ServiceProvider) *RegisterUserUseCase {
	return &RegisterUserUseCase{
		ServiceProvider: serviceProvider,
	}
}

func (u *RegisterUserUseCase) InitServices() {
	u.UserService = u.ServiceProvider.MakeService(config.GetDB(), "users")
}

func (u *RegisterUserUseCase) Invoke(params RegisterUserParams) (bool, error) {
	var count int64

	filterQuery := map[string]utils.EloquentQuery{
		"username": utils.GetExactMatchFilter(params.Body.Username),
	}

	u.UserService.Count(&count, params.Context, service.ServiceOption{
		Filter: filterQuery,
	})

	if count > 0 {
		fmt.Println("❌ Username already used")
		return false, errors.New("username already used")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(params.Body.Password), bcrypt.DefaultCost)
	params.Body.Password = string(hashedPassword)

	if err != nil {
		return false, errors.New("failed to generate bcrypt from password")
	}

	if err := u.UserService.InsertOne(params.Context, &params.Body, service.ServiceOption{}); err != nil {
		return false, err
	}

	return true, nil
}
