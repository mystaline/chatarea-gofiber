package dto

type EditRoomBody struct {
	Name string `json:"name" validate:"required"`
}

type CreateRoomBody struct {
	Type     string `json:"type"     validate:"required,oneof=group direct"`
	Name     string `json:"name"     validate:"required_if=Type group"`
	PersonID string `json:"personId" validate:"required_if=Type direct,uuid4"`
}

type CreateRoomResponse struct {
	Address string `json:"address"`
}

type RoomQuery struct {
	BaseQuery
	IsManaged string `query:"isManaged"`
}
