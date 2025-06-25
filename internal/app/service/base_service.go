package service

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/mystaline/chatarea-gofiber/internal/app/utils"
	"gorm.io/gorm"
)

type ServiceOption struct {
	Filter  map[string]utils.EloquentQuery
	Preload []string
}

type BaseService interface {
	FindMany(v interface{}, c *fiber.Ctx, option ServiceOption) error
	FindOne(v any, c *fiber.Ctx, option ServiceOption) error
	Count(v *int64, c *fiber.Ctx, option ServiceOption) error
	InsertOne(c *fiber.Ctx, body interface{}, option ServiceOption) error
	UpdateOne(v any, c *fiber.Ctx, body interface{}, option ServiceOption) error
	DeleteOne(c *fiber.Ctx, option ServiceOption) error
}

type BaseRepository struct {
	db        *gorm.DB
	tableName string
}

// FindMany implements BaseService.
func (g *BaseRepository) FindMany(
	v interface{},
	c *fiber.Ctx,
	option ServiceOption,
) error {
	res := g.db.Table(g.tableName)

	queried := utils.ApplyFilter(res, option.Filter)
	fmt.Println("queried", queried)

	if len(option.Preload) > 0 {
		for _, each := range option.Preload {
			queried = queried.Preload(each)
			fmt.Println("queried preload", queried)
		}
	}

	if err := queried.Find(v).Error; err != nil {
		return errors.New("failed to retrieve data")
	}
	fmt.Println("v", v)

	return nil
}

// FindOne implements BaseService.
func (g *BaseRepository) FindOne(
	v any,
	c *fiber.Ctx,
	option ServiceOption,
) error {
	res := g.db.Table(g.tableName)

	queried := utils.ApplyFilter(res, option.Filter)
	if queried.Error != nil {
		return errors.New("failed to apply filter")
	}

	if len(option.Preload) > 0 {
		for _, each := range option.Preload {
			queried.Preload(each)
		}
	}

	if queried.First(v).Error != nil {
		return errors.New("failed to retrieve data")
	}

	return nil
}

// Count implements BaseService.
func (g *BaseRepository) Count(
	v *int64,
	c *fiber.Ctx,
	option ServiceOption,
) error {
	res := g.db.Table(g.tableName)

	queried := utils.ApplyFilter(res, option.Filter)

	if queried.Count(v).Error != nil {
		return errors.New("failed to retrieve data")
	}

	return nil
}

// UpdateOne implements BaseService.
func (g *BaseRepository) InsertOne(
	c *fiber.Ctx,
	body interface{},
	option ServiceOption,
) error {
	fmt.Println("body", body)
	if err := g.db.Table(g.tableName).Create(body).Error; err != nil {
		return err
	}

	return nil
}

// UpdateOne implements BaseService.
func (g *BaseRepository) UpdateOne(
	v any,
	c *fiber.Ctx,
	body interface{},
	option ServiceOption,
) error {
	if err := g.FindOne(v, c, option); err != nil {
		return err
	}

	if res := g.db.Table(g.tableName).Model(v).Updates(body); res.Error != nil {
		return errors.New("failed to edit data")
	}

	return nil
}

// DeleteOne implements BaseService.
func (g *BaseRepository) DeleteOne(
	c *fiber.Ctx,
	option ServiceOption,
) error {
	var v any

	if err := g.FindOne(v, c, option); err != nil {
		return err
	}

	if res := g.db.Table(g.tableName).Delete(v); res.Error != nil {
		return errors.New("failed to delete account")
	}

	return nil
}

func MakeService(db *gorm.DB, tableName string) BaseService {
	return &BaseRepository{
		db:        db,
		tableName: tableName,
	}
}
