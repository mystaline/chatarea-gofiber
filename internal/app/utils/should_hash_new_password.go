package utils

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func ValidateNewPassword(
	oldPassword string,
	c *fiber.Ctx,
	newPassword string,
) (string, error) {
	if newPassword == "" {
		return "", errors.New("error not changed")
	}

	if bcrypt.CompareHashAndPassword([]byte(oldPassword), []byte(newPassword)) == nil {
		return "", errors.New("error not changed")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("failed to hash new password")
	}

	return string(hashedPassword), nil
}
