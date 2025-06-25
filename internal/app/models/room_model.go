package models

import (
	"time"

	"github.com/google/uuid"
)

type Room struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Name      string    `gorm:"size:255;not null"                              json:"name"`
	Address   string    `gorm:"size:255;unique;not null"                       json:"address"`
	Type      string
	CreatorID uuid.UUID `                                                      json:"creatorId"`
	Creator   User      `gorm:"foreignKey:CreatorID"                           json:"creator"`
	CreatedAt time.Time `                                                      json:"createdAt"`
	UpdatedAt time.Time `                                                      json:"updatedAt"`
}
