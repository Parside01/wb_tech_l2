package service

import (
	"context"
	"github.com/Parside01/dev11/internal/entity"
	"github.com/Parside01/dev11/internal/repository"
)

type UserService interface {
	CreateUser(ctx context.Context, userId int) (*entity.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) CreateUser(ctx context.Context, userId int) (*entity.User, error) {
	return s.repo.CreateUser(ctx, userId)
}
