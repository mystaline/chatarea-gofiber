package e2e

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/mystaline/chatarea-gofiber/internal/app/dto"
	"github.com/mystaline/chatarea-gofiber/internal/app/models"
	"github.com/mystaline/chatarea-gofiber/internal/config"
	router "github.com/mystaline/chatarea-gofiber/internal/router"
)

var app *fiber.App

func TestMain(m *testing.M) {
	_ = godotenv.Load(".env.test") // Load test env
	config.ConnectDB()             // Connect to test DB

	err := config.GetDB().AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("❌ Failed to migrate schema: %v", err)
	}

	app = fiber.New()
	router.SetupRoutes(app)

	code := m.Run()
	os.Exit(code)
}

func TestRegister_Success(t *testing.T) {
	config.GetDB().Exec("DELETE FROM users WHERE username = 'testuser'")

	payload := dto.RegisterBody{
		Name:     "Test User",
		Username: "testuser",
		Password: "secret123",
	}
	body, _ := json.Marshal(payload)

	req := httptest.NewRequest("POST", "/api/v1/auth/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected 201 Created, got %d", resp.StatusCode)
	}
}

func TestLogin_Success(t *testing.T) {
	body := map[string]string{
		"username": "testuser",
		"password": "secret123",
	}
	payload, _ := json.Marshal(body)

	req := httptest.NewRequest("POST", "/api/v1/auth/login", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected 200 OK, got %d", resp.StatusCode)
	}
}
