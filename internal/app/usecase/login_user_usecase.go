package usecase

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/mystaline/chatarea-gofiber/internal/app/dto"
	"github.com/mystaline/chatarea-gofiber/internal/app/provider"
	"github.com/mystaline/chatarea-gofiber/internal/app/service"
	"github.com/mystaline/chatarea-gofiber/internal/config"
	"golang.org/x/crypto/bcrypt"
)

type LoginUserParams struct {
	Context  *fiber.Ctx
	Options  service.ServiceOption
	Body     dto.LoginBody
	Response *dto.LoginResponse
}

type LoginUserUseCase struct {
	UserService     service.BaseService
	ServiceProvider provider.ServiceProvider
}

func MakeLoginUserUseCase(serviceProvider provider.ServiceProvider) *LoginUserUseCase {
	return &LoginUserUseCase{
		ServiceProvider: serviceProvider,
	}
}

func (u *LoginUserUseCase) InitServices() {
	u.UserService = u.ServiceProvider.MakeService(config.GetDB(), "users")
}

func (u *LoginUserUseCase) Invoke(params LoginUserParams) (bool, error) {
	var user dto.LoginInfo

	if err := u.UserService.FindOne(&user, params.Context, params.Options); err != nil {
		fmt.Println("❌ User not found")
		return false, errors.New("username not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(params.Body.Password)); err != nil {
		return false, errors.New("invalid credentials")
	}

	token, err := config.GenerateJWT(user.ID)
	if err != nil {
		return false, errors.New("failed to generate token")
	}

	params.Response.Token = token

	return true, nil
}
