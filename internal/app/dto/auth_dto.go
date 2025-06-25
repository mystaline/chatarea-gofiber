package dto

import "github.com/google/uuid"

type LoginBody struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginInfo struct {
	ID       uuid.UUID
	Username string
	Password string
}

type LoginResponse struct {
	Token string
}

type RegisterBody struct {
	Name     string `json:"name"     validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
