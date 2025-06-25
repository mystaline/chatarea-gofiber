package usecase

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mystaline/chatarea-gofiber/internal/app/dto"
	"github.com/mystaline/chatarea-gofiber/internal/app/models"
	"github.com/mystaline/chatarea-gofiber/internal/app/provider"
	"github.com/mystaline/chatarea-gofiber/internal/app/service"
	"github.com/mystaline/chatarea-gofiber/internal/app/utils"
	"github.com/mystaline/chatarea-gofiber/internal/config"
)

type EditRoomParams struct {
	Context *fiber.Ctx
	Body    dto.EditRoomBody
}

type EditRoomUseCase struct {
	RoomService     service.BaseService
	ServiceProvider provider.ServiceProvider
}

func MakeEditRoomUseCase(serviceProvider provider.ServiceProvider) *EditRoomUseCase {
	return &EditRoomUseCase{
		ServiceProvider: serviceProvider,
	}
}

func (u *EditRoomUseCase) InitServices() {
	u.RoomService = u.ServiceProvider.MakeService(config.GetDB(), "rooms")
}

func (u *EditRoomUseCase) Invoke(params EditRoomParams) (bool, error) {
	roomId := params.Context.Params("id")
	var room models.Room

	if params.Body.Name != "" {
		room.Name = params.Body.Name

		if err := u.RoomService.UpdateOne(nil, params.Context, room, service.ServiceOption{
			Filter: map[string]utils.EloquentQuery{
				"id": utils.GetExactMatchFilter(roomId),
			},
		}); err != nil {
			return false, err
		}
	}

	return true, nil
}
