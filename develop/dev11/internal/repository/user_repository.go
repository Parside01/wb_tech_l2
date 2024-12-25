package repository

import (
	"context"
	"errors"
	"github.com/Parside01/dev11/internal/entity"
	"sync"
	"time"
)

var (
	ErrUserNotFound  = errors.New("user not found")
	ErrEventNotFound = errors.New("event not found")
)

type UserRepository interface {
	CreateUser(ctx context.Context, userId int) (*entity.User, error)
	AddEventToUser(ctx context.Context, userId int, event *entity.Event) error
	UpdateEvent(ctx context.Context, userId int, event *entity.Event) error
	DeleteEvent(ctx context.Context, userId int, event *entity.Event) error
	GetEventsByTimeInterval(ctx context.Context, userId int, begin, end time.Time) ([]*entity.Event, error)
}

type userRepository struct {
	mutex sync.RWMutex // data race нам не нужен.)
	users map[int]*entity.User
}

func NewUserRepository() UserRepository {
	return &userRepository{
		mutex: sync.RWMutex{},
		users: make(map[int]*entity.User),
	}
}

func (r *userRepository) CreateUser(ctx context.Context, userId int) (*entity.User, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		r.mutex.Lock()
		defer r.mutex.Unlock()

		user := &entity.User{
			ID:     userId,
			Events: make(map[time.Time][]*entity.Event),
		}
		r.users[userId] = user
		return user, nil
	}
}

func (r *userRepository) AddEventToUser(ctx context.Context, userId int, event *entity.Event) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		r.mutex.Lock()
		defer r.mutex.Unlock()

		user, ok := r.users[userId]
		if !ok {
			return ErrUserNotFound
		}

		user.Events[event.Date] = append(user.Events[event.Date], event)
		return nil
	}
}

// UpdateEvent Важно, чтобы @event был полностью заполнен всей нужной информации о событии.
func (r *userRepository) UpdateEvent(ctx context.Context, userId int, event *entity.Event) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		// На самом деле проблема, если данных будет много, то mutex будет захвачен долго. Проблема in-memory.
		r.mutex.Lock()
		defer r.mutex.Unlock()

		user, ok := r.users[userId]
		if !ok {
			return ErrUserNotFound
		}

		events, ok := user.Events[event.Date]
		if !ok {
			return ErrEventNotFound
		}

		found := false
		for i := range events {
			if events[i].Title == event.Title {
				user.Events[event.Date][i] = event
				found = true
			}
		}
		if !found {
			return ErrEventNotFound
		}
		return nil
	}
}

func (r *userRepository) DeleteEvent(ctx context.Context, userId int, event *entity.Event) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		r.mutex.Lock()
		defer r.mutex.Unlock()

		user, ok := r.users[userId]
		if !ok {
			return ErrUserNotFound
		}

		events, ok := user.Events[event.Date]
		if !ok {
			return ErrEventNotFound
		}

		for i := range events {
			if events[i].Title == event.Title {
				user.Events[event.Date] = append(events[:i], events[i+1:]...)
				return nil
			}
		}
		return nil
	}
}

func (r *userRepository) GetEventsByTimeInterval(ctx context.Context, userId int, begin, end time.Time) ([]*entity.Event, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		r.mutex.RLock()
		defer r.mutex.RUnlock()

		user, exists := r.users[userId]
		if !exists {
			return nil, ErrUserNotFound
		}

		var result []*entity.Event
		for date, events := range user.Events {
			if date.After(begin) && date.Before(end) {
				result = append(result, events...)
			}
		}

		return result, nil
	}
}
