package service

import (
	"context"
	"todo/model"
	"todo/repository"
)

type Service interface {
	Create(ctx context.Context, name, desription, username string) (bool, error)
	GetAll(ctx context.Context) ([]model.Todo, error)
	GetByID(ctx context.Context, id string) (model.Todo, error)
	Update(ctx context.Context, todoUpdated model.Todo, id string) (bool, error)
	Delete(ctx context.Context, id string) (bool, error)
}

type service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) Create(ctx context.Context, name, desription, username string) (bool, error) {
	var todo = model.Todo{
		Name:        name,
		Description: desription,
		Username:    username,
	}
	created, err := s.repository.Create(ctx, todo)
	if err != nil {
		return false, err
	}
	return created, nil
}

func (s *service) GetAll(ctx context.Context) ([]model.Todo, error) {
	todos, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (s *service) GetByID(ctx context.Context, id string) (model.Todo, error) {
	todo, err := s.repository.GetByID(ctx, id)
	if err != nil {
		return todo, err
	}
	return todo, nil
}

func (s *service) Update(ctx context.Context, todoUpdated model.Todo, id string) (bool, error) {
	updated, err := s.repository.Update(ctx, todoUpdated, id)
	if err != nil {
		return false, err
	}
	return updated, nil
}

func (s *service) Delete(ctx context.Context, id string) (bool, error) {
	deleted, err := s.repository.Delete(ctx, id)
	if err != nil {
		return false, err
	}
	return deleted, nil
}
