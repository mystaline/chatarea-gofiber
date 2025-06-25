package usecase

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mystaline/chatarea-gofiber/internal/app/models"
	"github.com/mystaline/chatarea-gofiber/internal/app/provider"
	"github.com/mystaline/chatarea-gofiber/internal/app/service"
	"github.com/mystaline/chatarea-gofiber/internal/app/utils"
	"github.com/mystaline/chatarea-gofiber/internal/config"
)

type GetRoomInfoParams struct {
	Context  *fiber.Ctx
	Response *models.Room
}

type GetRoomInfoUseCase struct {
	RoomService     service.BaseService
	UserService     service.BaseService
	UserRoomService service.BaseService
	ServiceProvider provider.ServiceProvider
}

func MakeGetRoomInfoUseCase(serviceProvider provider.ServiceProvider) *GetRoomInfoUseCase {
	return &GetRoomInfoUseCase{
		ServiceProvider: serviceProvider,
	}
}

func (u *GetRoomInfoUseCase) InitServices() {
	u.RoomService = u.ServiceProvider.MakeService(config.GetDB(), "rooms")
	u.UserService = u.ServiceProvider.MakeService(config.GetDB(), "users")
	u.UserRoomService = u.ServiceProvider.MakeService(config.GetDB(), "room_members")
}

func (u *GetRoomInfoUseCase) Invoke(params GetRoomInfoParams) (bool, error) {
	roomId := params.Context.Params("id")

	if err := u.RoomService.FindOne(params.Response, params.Context, service.ServiceOption{
		Filter: map[string]utils.EloquentQuery{
			"id": utils.GetExactMatchFilter(roomId),
		},
	}); err != nil {
		return false, err
	}

	return true, nil
}
