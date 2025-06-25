package controllers

// import (
// 	"bytes"
// 	"encoding/json"
// 	"errors"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/google/uuid"
// 	"github.com/mystaline/chatarea-gofiber/internal/app/dto"
// 	"github.com/mystaline/chatarea-gofiber/internal/app/mocks"
// 	"github.com/mystaline/chatarea-gofiber/internal/app/models"
// 	"github.com/mystaline/chatarea-gofiber/internal/app/service"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// func Test_Me_NotFound(t *testing.T) {
// 	app := setupTestApp()
// 	mockService := mocks.CreateMockService[models.User]().(*mocks.MockBaseRepository[models.User])
// 	service.UserService = mockService

// 	mockService.On("FindOne", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("user not found"))

// 	req := httptest.NewRequest("GET", "/api/v1/me/profile", nil)
// 	req.Header.Set("Content-Type", "application/json")
// 	resp, _ := app.Test(req)

// 	assert.Equal(t, 404, resp.StatusCode)
// 	mockService.AssertExpectations(t)
// }

// func Test_Me_Success(t *testing.T) {
// 	app := setupTestApp()
// 	mockService := mocks.CreateMockService[models.User]().(*mocks.MockBaseRepository[models.User])
// 	service.UserService = mockService

// 	mockService.On("FindOne", mock.Anything, mock.Anything, mock.Anything).Return(nil).Run(func(args mock.Arguments) {
// 		ptr := args.Get(0).(*models.User)
// 		*ptr = models.User{
// 			ID:   uuid.New(),
// 			Name: "Test User",
// 		}
// 	})

// 	req := httptest.NewRequest("GET", "/api/v1/me/profile", nil)
// 	req.Header.Set("Content-Type", "application/json")
// 	resp, _ := app.Test(req)

// 	assert.Equal(t, 200, resp.StatusCode)
// 	mockService.AssertExpectations(t)
// }

// func Test_EditProfile_InvalidBody(t *testing.T) {
// 	app := setupTestApp()
// 	mockService := mocks.CreateMockService[models.User]().(*mocks.MockBaseRepository[models.User])
// 	service.UserService = mockService

// 	req := httptest.NewRequest("PUT", "/api/v1/me/profile", bytes.NewBufferString("not json"))
// 	req.Header.Set("Content-Type", "application/json")
// 	resp, _ := app.Test(req)

// 	assert.Equal(t, 400, resp.StatusCode)
// 	mockService.AssertExpectations(t)
// }

// func Test_EditProfile_NotFound(t *testing.T) {
// 	app := setupTestApp()
// 	mockService := mocks.CreateMockService[models.User]().(*mocks.MockBaseRepository[models.User])
// 	service.UserService = mockService

// 	payload := dto.EditProfileBody{
// 		Name: "Test User 2",
// 	}
// 	body, _ := json.Marshal(payload)

// 	mockService.On("UpdateOne", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
// 		Return(errors.New("retrieve data"))

// 	req := httptest.NewRequest("PUT", "/api/v1/me/profile", bytes.NewBuffer(body))
// 	req.Header.Set("Content-Type", "application/json")
// 	resp, _ := app.Test(req)

// 	assert.Equal(t, 404, resp.StatusCode)
// 	mockService.AssertExpectations(t)
// }

// func Test_EditProfile_InternalServerError(t *testing.T) {
// 	app := setupTestApp()
// 	mockService := mocks.CreateMockService[models.User]().(*mocks.MockBaseRepository[models.User])
// 	service.UserService = mockService

// 	payload := dto.EditProfileBody{
// 		Name: "Test User 2",
// 	}
// 	body, _ := json.Marshal(payload)

// 	mockService.On("UpdateOne", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
// 		Return(errors.New("edit data"))

// 	req := httptest.NewRequest("PUT", "/api/v1/me/profile", bytes.NewBuffer(body))
// 	req.Header.Set("Content-Type", "application/json")
// 	resp, _ := app.Test(req)

// 	assert.Equal(t, 500, resp.StatusCode)
// 	mockService.AssertExpectations(t)
// }

// func Test_EditProfile_Success(t *testing.T) {
// 	app := setupTestApp()
// 	mockService := mocks.CreateMockService[models.User]().(*mocks.MockBaseRepository[models.User])
// 	service.UserService = mockService

// 	payload := dto.EditProfileBody{
// 		Name: "Test User 2",
// 	}
// 	body, _ := json.Marshal(payload)

// 	mockService.On("UpdateOne", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

// 	req := httptest.NewRequest("PUT", "/api/v1/me/profile", bytes.NewBuffer(body))
// 	req.Header.Set("Content-Type", "application/json")
// 	resp, _ := app.Test(req)

// 	assert.Equal(t, 200, resp.StatusCode)
// 	mockService.AssertExpectations(t)
// }
