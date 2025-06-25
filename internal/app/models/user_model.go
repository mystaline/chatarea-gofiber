package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID             uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Name           string    `gorm:"size:255;not null"                              json:"name"`
	Username       string    `gorm:"size:255;unique;not null"                       json:"username"`
	Password       string    `gorm:"size:255;not null"                              json:"-"`
	Timezone       *string   `gorm:"size:255"                                       json:"timezone,omitempty"`
	RememberToken  *string   `gorm:"size:100"                                       json:"-"`
	CreatedAt      time.Time `                                                      json:"createdAt"`
	UpdatedAt      time.Time `                                                      json:"updatedAt"`
	ProfilePicture *string   `gorm:"size:255"                                       json:"profilePicture,omitempty"`
}

type SimpleUser struct {
	ID   uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Name string    `gorm:"size:255;not null"                              json:"name"`
}
