package usecase

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mystaline/chatarea-gofiber/internal/app/models"
	"github.com/mystaline/chatarea-gofiber/internal/app/provider"
	"github.com/mystaline/chatarea-gofiber/internal/app/service"
	"github.com/mystaline/chatarea-gofiber/internal/config"
)

type GetMyRoomsParams struct {
	Context  *fiber.Ctx
	Options  service.ServiceOption
	Response *[]models.UserRoom
}

type GetMyRoomsUseCase struct {
	UserRoomService service.BaseService
	ServiceProvider provider.ServiceProvider
}

func MakeGetMyRoomsUseCase(serviceProvider provider.ServiceProvider) *GetMyRoomsUseCase {
	return &GetMyRoomsUseCase{
		ServiceProvider: serviceProvider,
	}
}

func (u *GetMyRoomsUseCase) InitServices() {
	u.UserRoomService = u.ServiceProvider.MakeService(config.GetDB(), "room_members")
}

func (u *GetMyRoomsUseCase) Invoke(params GetMyRoomsParams) (bool, error) {
	if err := u.UserRoomService.FindMany(params.Response, params.Context, params.Options); err != nil {
		return false, err
	}

	return true, nil
}
