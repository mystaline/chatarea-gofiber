package dto

import (
	"time"

	"github.com/google/uuid"
)

type UserQuery struct {
	BaseQuery
	ID       uuid.UUID `query:"id"`
	Name     string    `query:"name"`
	Username string    `query:"username"`
}

type EditProfileBody struct {
	Name           string `json:"name"`
	Username       string `json:"username"`
	Password       string `json:"password"`
	ProfilePicture string `json:"profilePicture"`
}

type GetMyProfileResponse struct {
	ID             uuid.UUID `json:"id"`
	Name           string    `json:"name"`
	Username       string    `json:"username"`
	Password       string    `json:"-"`
	Timezone       *string   `json:"timezone,omitempty"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	ProfilePicture *string   `json:"profilePicture,omitempty"`
}

type GetUserListResponse struct {
	ID             uuid.UUID `json:"id"`
	Name           string    `json:"name"`
	Username       string    `json:"username"`
	CreatedAt      time.Time `json:"createdAt"`
	ProfilePicture *string   `json:"profilePicture,omitempty"`
}
