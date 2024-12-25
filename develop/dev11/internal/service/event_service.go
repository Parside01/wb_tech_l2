package service

import (
	"context"
	"github.com/Parside01/dev11/internal/entity"
	"github.com/Parside01/dev11/internal/repository"
	"time"
)

type EventService interface {
	CreateEvent(ctx context.Context, userId int, event *entity.Event) error
	UpdateEvent(ctx context.Context, userId int, event *entity.Event) error
	DeleteEvent(ctx context.Context, userId int, event *entity.Event) error
	EventsByDay(ctx context.Context, userId int, currTime time.Time) ([]*entity.Event, error)
	EventsByWeek(ctx context.Context, userId int, currTime time.Time) ([]*entity.Event, error)
	EventsByMonth(ctx context.Context, userId int, currTime time.Time) ([]*entity.Event, error)
}

type eventService struct {
	repo repository.UserRepository
}

func NewEventService(repo repository.UserRepository) EventService {
	return &eventService{
		repo: repo,
	}
}

func (s *eventService) CreateEvent(ctx context.Context, userId int, event *entity.Event) error {
	return s.repo.AddEventToUser(ctx, userId, event)
}

func (s *eventService) UpdateEvent(ctx context.Context, userId int, event *entity.Event) error {
	return s.repo.UpdateEvent(ctx, userId, event)
}

func (s *eventService) DeleteEvent(ctx context.Context, userId int, event *entity.Event) error {
	return s.repo.DeleteEvent(ctx, userId, event)
}

// Не очень понимаю какие могут быть ошибки в бизнес логике на этом этапе, но ладно.
func (s *eventService) EventsByDay(ctx context.Context, userId int, currTime time.Time) ([]*entity.Event, error) {
	start := currTime.Truncate(24 * time.Hour)
	end := start.Add(24 * time.Hour)

	events, err := s.repo.GetEventsByTimeInterval(ctx, userId, start, end)
	if err != nil {
		return nil, err
	}
	return events, nil
}
func (s *eventService) EventsByWeek(ctx context.Context, userId int, currTime time.Time) ([]*entity.Event, error) {
	start := currTime.Truncate(24*time.Hour).AddDate(0, 0, -int(currTime.Weekday()))
	end := start.AddDate(0, 0, 7)

	return s.repo.GetEventsByTimeInterval(ctx, userId, start, end)
}
func (s *eventService) EventsByMonth(ctx context.Context, userId int, currTime time.Time) ([]*entity.Event, error) {
	start := time.Date(currTime.Year(), currTime.Month(), 1, 0, 0, 0, 0, currTime.Location())
	end := start.AddDate(0, 1, 0)

	return s.repo.GetEventsByTimeInterval(ctx, userId, start, end)
}
