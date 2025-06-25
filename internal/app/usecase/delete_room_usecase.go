package usecase

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mystaline/chatarea-gofiber/internal/app/provider"
	"github.com/mystaline/chatarea-gofiber/internal/app/service"
	"github.com/mystaline/chatarea-gofiber/internal/app/utils"
	"github.com/mystaline/chatarea-gofiber/internal/config"
)

type DeleteRoomParams struct {
	Context *fiber.Ctx
}

type DeleteRoomUseCase struct {
	RoomService     service.BaseService
	ServiceProvider provider.ServiceProvider
}

func MakeDeleteRoomUseCase(serviceProvider provider.ServiceProvider) *DeleteRoomUseCase {
	return &DeleteRoomUseCase{
		ServiceProvider: serviceProvider,
	}
}

func (u *DeleteRoomUseCase) InitServices() {
	u.RoomService = u.ServiceProvider.MakeService(config.GetDB(), "rooms")
}

func (u *DeleteRoomUseCase) Invoke(params DeleteRoomParams) (bool, error) {
	roomId := params.Context.Params("id")

	if err := u.RoomService.DeleteOne(params.Context, service.ServiceOption{
		Filter: map[string]utils.EloquentQuery{
			"id": utils.GetExactMatchFilter(roomId),
		},
	}); err != nil {
		return false, err
	}

	return true, nil
}
