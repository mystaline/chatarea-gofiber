package middleware

import (
	"fmt"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/mystaline/chatarea-gofiber/internal/app/dto"
	"github.com/mystaline/chatarea-gofiber/internal/app/provider"
	"github.com/mystaline/chatarea-gofiber/internal/app/service"
	"github.com/mystaline/chatarea-gofiber/internal/app/usecase"
	"github.com/mystaline/chatarea-gofiber/internal/app/utils"
)

func ValidateJWT() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get(fiber.HeaderAuthorization)

		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			return c.Status(fiber.ErrUnauthorized.Code).JSON(fiber.Map{
				"status":  fiber.ErrUnauthorized.Code,
				"message": "Invalid Authorization header",
			})
		}

		splittedToken := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(splittedToken, func(t *jwt.Token) (any, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}

			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			return c.Status(fiber.ErrUnauthorized.Code).JSON(fiber.Map{
				"status":  fiber.ErrUnauthorized.Code,
				"message": "Invalid or expired token",
			})
		}

		claims := token.Claims.(jwt.MapClaims)

		user := &dto.GetMyProfileResponse{}
		useCase := usecase.MakeGetProfileUseCase(&provider.MainServiceProvider{})
		useCase.InitServices()

		filterQuery := map[string]utils.EloquentQuery{
			"id": utils.GetExactMatchFilter(claims["id"]),
		}

		if _, err := useCase.Invoke(usecase.GetProfileParams{
			Context:  c,
			Response: user,
			Options:  service.ServiceOption{Filter: filterQuery},
		}); err != nil {
			return c.Status(fiber.ErrUnauthorized.Code).JSON(fiber.Map{
				"status":  fiber.ErrUnauthorized.Code,
				"message": "Account not found or already deleted",
			})
		}
		fmt.Println("auth user", user)

		c.Locals("user", user)

		return c.Next()
	}
}

func ValidateCookie() fiber.Handler {
	panic("not implemented")
}
