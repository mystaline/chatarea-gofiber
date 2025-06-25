package usecase

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	nanoid "github.com/matoous/go-nanoid"
	"github.com/mystaline/chatarea-gofiber/internal/app/dto"
	"github.com/mystaline/chatarea-gofiber/internal/app/models"
	"github.com/mystaline/chatarea-gofiber/internal/app/provider"
	"github.com/mystaline/chatarea-gofiber/internal/app/service"
	"github.com/mystaline/chatarea-gofiber/internal/app/utils"
	"github.com/mystaline/chatarea-gofiber/internal/config"
)

type CreateRoomParams struct {
	Context  *fiber.Ctx
	Body     dto.CreateRoomBody
	Response dto.CreateRoomResponse
}

type CreateRoomUseCase struct {
	RoomService     service.BaseService
	UserService     service.BaseService
	UserRoomService service.BaseService
	ServiceProvider provider.ServiceProvider
}

func MakeCreateRoomUseCase(serviceProvider provider.ServiceProvider) *CreateRoomUseCase {
	return &CreateRoomUseCase{
		ServiceProvider: serviceProvider,
	}
}

func (u *CreateRoomUseCase) InitServices() {
	u.RoomService = u.ServiceProvider.MakeService(config.GetDB(), "rooms")
	u.UserService = u.ServiceProvider.MakeService(config.GetDB(), "users")
	u.UserRoomService = u.ServiceProvider.MakeService(config.GetDB(), "room_members")
}

func (u *CreateRoomUseCase) Invoke(params CreateRoomParams) (bool, error) {
	user, ok := params.Context.Locals("user").(*dto.GetMyProfileResponse)
	if !ok {
		return false, errors.New("invalid user context")
	}

	newRoom := models.Room{
		ID:        uuid.New(),
		Type:      params.Body.Type,
		Address:   nanoid.MustGenerate("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", 10),
		CreatorID: user.ID,
	}

	if params.Body.PersonID != "" {
		if err := u.RoomService.InsertOne(params.Context, &newRoom, service.ServiceOption{}); err != nil {
			return false, err
		}

		var targetPerson models.SimpleUser
		if err := u.UserService.FindOne(targetPerson, params.Context, service.ServiceOption{
			Filter: map[string]utils.EloquentQuery{
				"id": utils.GetExactMatchFilter(params.Body.PersonID),
			},
		}); err != nil {
			return false, errors.New("user not found")
		}

		targetPersonRoom := models.UserRoom{
			ID:     uuid.New(),
			UserID: targetPerson.ID,
			RoomID: newRoom.ID,
		}

		if err := u.UserRoomService.InsertOne(params.Context, &targetPersonRoom, service.ServiceOption{}); err != nil {
			return false, err
		}
	} else if params.Body.Name != "" {
		newRoom.Name = params.Body.Name

		if err := u.RoomService.InsertOne(params.Context, &newRoom, service.ServiceOption{}); err != nil {
			return false, err
		}
	}

	userRoom := models.UserRoom{
		ID:     uuid.New(),
		UserID: user.ID,
		RoomID: newRoom.ID,
	}

	if err := u.UserRoomService.InsertOne(params.Context, &userRoom, service.ServiceOption{}); err != nil {
		return false, err
	}

	params.Response.Address = newRoom.Address

	return true, nil
}
