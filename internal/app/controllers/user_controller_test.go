package controllers

// import (
// 	"errors"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/mystaline/chatarea-gofiber/internal/app/mocks"
// 	"github.com/mystaline/chatarea-gofiber/internal/app/models"
// 	"github.com/mystaline/chatarea-gofiber/internal/app/service"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// func Test_GetUsers_InvalidQuery(t *testing.T) {
// 	app := setupTestApp()
// 	mockService := mocks.CreateMockService[models.User]().(*mocks.MockBaseRepository[models.User])
// 	service.UserService = mockService

// 	req := httptest.NewRequest("GET", "/ping?id={12}", nil)
// 	req.Header.Set("Content-Type", "application/json")
// 	resp, _ := app.Test(req)

// 	assert.Equal(t, 400, resp.StatusCode)
// 	mockService.AssertExpectations(t)
// }

// func Test_GetUsers_InternalServerError(t *testing.T) {
// 	app := setupTestApp()
// 	mockService := mocks.CreateMockService[models.User]().(*mocks.MockBaseRepository[models.User])
// 	service.UserService = mockService

// 	mockService.On("FindMany", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("internal server error"))

// 	req := httptest.NewRequest("GET", "/ping", nil)
// 	req.Header.Set("Content-Type", "application/json")
// 	resp, _ := app.Test(req)

// 	assert.Equal(t, 500, resp.StatusCode)
// 	mockService.AssertExpectations(t)
// }

// func Test_GetUsers_Success(t *testing.T) {
// 	app := setupTestApp()
// 	mockService := mocks.CreateMockService[models.User]().(*mocks.MockBaseRepository[models.User])
// 	service.UserService = mockService

// 	mockService.On("FindMany", mock.Anything, mock.Anything, mock.Anything).Return(nil)

// 	req := httptest.NewRequest("GET", "/ping", nil)
// 	req.Header.Set("Content-Type", "application/json")
// 	resp, _ := app.Test(req)

// 	assert.Equal(t, 200, resp.StatusCode)
// 	mockService.AssertExpectations(t)
// }
