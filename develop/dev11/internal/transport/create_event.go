package transport

import (
	"github.com/Parside01/dev11/internal/repository"
	"github.com/Parside01/dev11/internal/service"
	"net/http"
)

type CreateEventHandler interface {
	CreateEvenHandler(w http.ResponseWriter, r *http.Request)
}

type createEventHandler struct {
	service service.EventService
}

func NewCreateHandler(repo *repository.UserRepository) CreateEventHandler {
	return &createEventHandler{
		repo: repo,
	}
}

func (h *createEventHandler) CreateEvenHandler(w http.ResponseWriter, r *http.Request) {
	data, err := PostDataFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
