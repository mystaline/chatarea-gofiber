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
	Joins   []string
	Select  []string
}

type BaseService interface {
	SQLBuilder(option ServiceOption) *gorm.DB

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

func (g *BaseRepository) SQLBuilder(option ServiceOption) *gorm.DB {
	res := g.db.Table(g.tableName)

	queried := utils.ApplyFilter(res, option.Filter)

	if len(option.Joins) > 0 {
		for _, each := range option.Joins {
			queried = queried.Joins(each)
		}
	}

	fmt.Println("option.Select", option.Select)
	if len(option.Select) > 0 {
		queried = queried.Select(option.Select)
	}

	if len(option.Preload) > 0 {
		for _, each := range option.Preload {
			queried = queried.Preload(each)
			fmt.Println("queried preload", queried)
		}
	}

	return queried
}

// FindMany implements BaseService.
func (g *BaseRepository) FindMany(
	v interface{},
	c *fiber.Ctx,
	option ServiceOption,
) error {
	sqlBuilt := g.SQLBuilder(option)

	if err := sqlBuilt.Find(v).Error; err != nil {
		return errors.New("failed to retrieve data")
	}
	fmt.Printf("%+v\n", v)

	return nil
}

// FindOne implements BaseService.
func (g *BaseRepository) FindOne(
	v any,
	c *fiber.Ctx,
	option ServiceOption,
) error {
	sqlBuilt := g.SQLBuilder(option)

	if sqlBuilt.First(v).Error != nil {
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
	sqlBuilt := g.SQLBuilder(option)

	if sqlBuilt.Count(v).Error != nil {
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
	sqlBuilt := g.SQLBuilder(option)

	if err := sqlBuilt.Create(body).Error; err != nil {
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

	if sqlBuilt := g.SQLBuilder(option).Model(v).Updates(body); sqlBuilt.Error != nil {
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
		db:        db.Debug(),
		tableName: tableName,
	}
}
