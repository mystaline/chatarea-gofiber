package provider

import (
	"github.com/mystaline/chatarea-gofiber/internal/app/service"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type ServiceProvider interface {
	MakeService(db *gorm.DB, tableName string) service.BaseService
}

type MainServiceProvider struct{}

func (m *MainServiceProvider) MakeService(db *gorm.DB, tableName string) service.BaseService {
	return service.MakeService(db, tableName)
}

type MockServiceProvider struct {
	mock.Mock
}

func (m *MockServiceProvider) MakeService(db *gorm.DB, tableName string) service.BaseService {
	args := m.Called(db, tableName)
	return args.Error(0).(service.BaseService)
}
