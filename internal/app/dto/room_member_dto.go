package dto

import (
	"time"

	"github.com/google/uuid"
)

type MyRoom struct {
	ID          uuid.UUID `gorm:"alias:room_members.id;column:id"           json:"id"`
	RoomID      uuid.UUID `gorm:"alias:room_members.room_id;column:room_id" json:"roomId"`
	Name        string    `gorm:"alias:rooms.name;column:name"              json:"name"`
	Address     string    `gorm:"alias:rooms.address;column:address"        json:"address"`
	Type        string    `gorm:"alias:rooms.type;column:type"              json:"type"`
	CreatorID   uuid.UUID `gorm:"alias:users.id;column:creator_id"          json:"creatorId"`
	CreatorName string    `gorm:"alias:users.name;column:creator_name"      json:"creatorName"`
	CreatedAt   time.Time `                                                 json:"createdAt"`
	UpdatedAt   time.Time `                                                 json:"updatedAt"`
}

type Creator struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type GetMyRoomsResponse struct {
	ID        uuid.UUID `json:"id"`
	RoomID    uuid.UUID `json:"roomId"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	Type      string    `json:"type"`
	Creator   Creator   `json:"creator"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
