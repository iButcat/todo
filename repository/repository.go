package repository

import (
	"context"
	"todo/model"

	"gorm.io/gorm"
)

type Repository interface {
	Create(ctx context.Context, todo model.Todo) (bool, error)
	GetAll(ctx context.Context) ([]model.Todo, error)
	GetByID(ctx context.Context, id string) (model.Todo, error)
	Update(ctx context.Context, todoUpdated model.Todo, id string) (bool, error)
	Delete(ctx context.Context, id string) (bool, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (repo *repository) Create(ctx context.Context, todo model.Todo) (bool, error) {
	if err := repo.db.AutoMigrate(&model.Todo{}); err != nil {
		return false, err
	}

	if err := repo.db.Create(&todo).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (repo *repository) GetAll(ctx context.Context) ([]model.Todo, error) {
	var todos []model.Todo
	if err := repo.db.Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}

func (repo *repository) GetByID(ctx context.Context, id string) (model.Todo, error) {
	var todo model.Todo
	if err := repo.db.Find(&todo, id).Error; err != nil {
		return todo, err
	}
	return todo, nil
}

func (repo *repository) Update(ctx context.Context, todoUpdated model.Todo, id string) (bool, error) {
	if err := repo.db.Model(&model.Todo{}).Where("id = ?", id).Updates(todoUpdated).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (repo *repository) Delete(ctx context.Context, id string) (bool, error) {
	if err := repo.db.Delete(&model.Todo{}, id).Error; err != nil {
		return false, err
	}
	return true, nil
}
