package models

import (
	"time"

	"github.com/google/uuid"
)

type RoomMember struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	UserID    uuid.UUID `                                                      json:"userId"`
	User      User      `gorm:"foreignKey:UserID"                              json:"user"`
	RoomID    uuid.UUID `                                                      json:"roomId"`
	Room      Room      `gorm:"foreignKey:RoomID"                              json:"room"`
	CreatedAt time.Time `                                                      json:"createdAt"`
	UpdatedAt time.Time `                                                      json:"updatedAt"`
}
