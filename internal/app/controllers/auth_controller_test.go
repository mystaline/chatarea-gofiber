package controllers

// import (
// 	"bytes"
// 	"encoding/json"
// 	"errors"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/mystaline/chatarea-gofiber/internal/app/dto"
// 	"github.com/mystaline/chatarea-gofiber/internal/app/mocks"
// 	"github.com/mystaline/chatarea-gofiber/internal/app/service"
// 	"github.com/stretchr/testify/assert"
// )

// func Test_Register_InvalidBody(t *testing.T) {
// 	app := setupTestApp()
// 	mockService := new(mocks.MockAuthRepository)
// 	service.AuthService = mockService

// 	req := httptest.NewRequest("POST", "/api/v1/auth/register", bytes.NewBufferString("not json"))
// 	req.Header.Set("Content-Type", "application/json")
// 	resp, _ := app.Test(req)

// 	assert.Equal(t, 400, resp.StatusCode)
// 	mockService.AssertExpectations(t)
// }

// func Test_Register_InvalidValidation(t *testing.T) {
// 	app := setupTestApp()
// 	mockService := new(mocks.MockAuthRepository)
// 	service.AuthService = mockService

// 	payload := dto.RegisterBody{}
// 	body, _ := json.Marshal(payload)

// 	req := httptest.NewRequest("POST", "/api/v1/auth/register", bytes.NewBuffer(body))
// 	req.Header.Set("Content-Type", "application/json")
// 	resp, _ := app.Test(req)

// 	assert.Equal(t, 400, resp.StatusCode)
// 	mockService.AssertExpectations(t)
// }

// func Test_Register_FailedBcrypt(t *testing.T) {
// 	app := setupTestApp()
// 	mockService := new(mocks.MockAuthRepository)
// 	service.AuthService = mockService

// 	mockService.On("RegisterUser", "Test User", "testuser", "secret123").
// 		Return(errors.New("failed to generate bcrypt from password"))

// 	payload := dto.RegisterBody{
// 		Name:     "Test User",
// 		Username: "testuser",
// 		Password: "secret123",
// 	}
// 	body, _ := json.Marshal(payload)

// 	req := httptest.NewRequest("POST", "/api/v1/auth/register", bytes.NewBuffer(body))
// 	req.Header.Set("Content-Type", "application/json")
// 	resp, _ := app.Test(req)

// 	assert.Equal(t, 500, resp.StatusCode)
// 	mockService.AssertExpectations(t)
// }

// func Test_Register_UsernameExist(t *testing.T) {
// 	app := setupTestApp()
// 	mockService := new(mocks.MockAuthRepository)
// 	service.AuthService = mockService

// 	mockService.On("RegisterUser", "Test User", "testuser", "secret123").
// 		Return(errors.New("username already used"))

// 	payload := dto.RegisterBody{
// 		Name:     "Test User",
// 		Username: "testuser",
// 		Password: "secret123",
// 	}
// 	body, _ := json.Marshal(payload)

// 	req := httptest.NewRequest("POST", "/api/v1/auth/register", bytes.NewBuffer(body))
// 	req.Header.Set("Content-Type", "application/json")
// 	resp, _ := app.Test(req)

// 	var response map[string]interface{}
// 	json.NewDecoder(resp.Body).Decode(&response)

// 	assert.Equal(t, 400, resp.StatusCode)
// 	assert.Contains(t, response["message"], "username already used")
// 	mockService.AssertExpectations(t)
// }

// func Test_Register_Success(t *testing.T) {
// 	app := setupTestApp()
// 	mockService := new(mocks.MockAuthRepository)
// 	service.AuthService = mockService

// 	mockService.On("RegisterUser", "Test User", "testuser", "secret123").Return(nil)

// 	payload := dto.RegisterBody{
// 		Name:     "Test User",
// 		Username: "testuser",
// 		Password: "secret123",
// 	}
// 	body, _ := json.Marshal(payload)

// 	req := httptest.NewRequest("POST", "/api/v1/auth/register", bytes.NewBuffer(body))
// 	req.Header.Set("Content-Type", "application/json")
// 	resp, _ := app.Test(req)

// 	assert.Equal(t, 201, resp.StatusCode)
// 	mockService.AssertExpectations(t)
// }

// func Test_Login_InvalidBody(t *testing.T) {
// 	app := setupTestApp()
// 	mockService := new(mocks.MockAuthRepository)
// 	service.AuthService = mockService

// 	req := httptest.NewRequest("POST", "/api/v1/auth/login", bytes.NewBufferString("not json"))
// 	req.Header.Set("Content-Type", "application/json")
// 	resp, _ := app.Test(req)

// 	assert.Equal(t, 400, resp.StatusCode)
// 	mockService.AssertExpectations(t)
// }

// func Test_Login_InvalidValidation(t *testing.T) {
// 	app := setupTestApp()
// 	mockService := new(mocks.MockAuthRepository)
// 	service.AuthService = mockService

// 	payload := dto.LoginBody{
// 		Username: "testuser",
// 	}
// 	body, _ := json.Marshal(payload)

// 	req := httptest.NewRequest("POST", "/api/v1/auth/login", bytes.NewBuffer(body))
// 	req.Header.Set("Content-Type", "application/json")
// 	resp, _ := app.Test(req)

// 	assert.Equal(t, 400, resp.StatusCode)
// 	mockService.AssertExpectations(t)
// }

// func Test_Login_UserNotFound(t *testing.T) {
// 	app := setupTestApp()
// 	mockService := new(mocks.MockAuthRepository)
// 	service.AuthService = mockService

// 	mockService.On("LoginUser", "testuser", "secret123").Return("", errors.New("user not found"))

// 	payload := dto.LoginBody{
// 		Username: "testuser",
// 		Password: "secret123",
// 	}
// 	body, _ := json.Marshal(payload)

// 	req := httptest.NewRequest("POST", "/api/v1/auth/login", bytes.NewBuffer(body))
// 	req.Header.Set("Content-Type", "application/json")
// 	resp, _ := app.Test(req)

// 	assert.Equal(t, 404, resp.StatusCode)
// 	mockService.AssertExpectations(t)
// }

// func Test_Login_InvalidCredentials(t *testing.T) {
// 	app := setupTestApp()
// 	mockService := new(mocks.MockAuthRepository)
// 	service.AuthService = mockService

// 	mockService.On("LoginUser", "testuser", "secret123").Return("", errors.New("invalid credentials"))

// 	payload := dto.LoginBody{
// 		Username: "testuser",
// 		Password: "secret123",
// 	}
// 	body, _ := json.Marshal(payload)

// 	req := httptest.NewRequest("POST", "/api/v1/auth/login", bytes.NewBuffer(body))
// 	req.Header.Set("Content-Type", "application/json")
// 	resp, _ := app.Test(req)

// 	assert.Equal(t, 401, resp.StatusCode)
// 	mockService.AssertExpectations(t)
// }

// func Test_Login_Success(t *testing.T) {
// 	app := setupTestApp()
// 	mockService := new(mocks.MockAuthRepository)
// 	service.AuthService = mockService

// 	mockService.On("LoginUser", "testuser", "secret123").Return("", nil)

// 	payload := dto.LoginBody{
// 		Username: "testuser",
// 		Password: "secret123",
// 	}
// 	body, _ := json.Marshal(payload)

// 	req := httptest.NewRequest("POST", "/api/v1/auth/login", bytes.NewBuffer(body))
// 	req.Header.Set("Content-Type", "application/json")
// 	resp, _ := app.Test(req)

// 	assert.Equal(t, 201, resp.StatusCode)
// 	mockService.AssertExpectations(t)
// }
