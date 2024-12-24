package service

import (
	"context"
	"github.com/Parside01/dev11/internal/entity"
	"github.com/Parside01/dev11/internal/repository"
)

type EventService interface {
	CreateEvent(ctx context.Context, userId int, event *entity.Event) error
	UpdateEvent(ctx context.Context, userId int, event *entity.Event) error
	DeleteEvent(ctx context.Context, userId int, event *entity.Event) error
}

type eventService struct {
	repo *repository.UserRepository
}

func NewEventService(repo *repository.UserRepository) EventService {
	return &eventService{
		repo: repo,
	}
}

func (s *eventService) CreateEvent(ctx context.Context, userId int, event *entity.Event) error {

}

func (s *eventService) UpdateEvent(ctx context.Context, userId int, event *entity.Event) error {

}

func (s *eventService) DeleteEvent(ctx context.Context, userId int, event *entity.Event) error {

}
