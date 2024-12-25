package transport

import (
	"github.com/Parside01/dev11/internal/service"
	"net/http"
)

type CreateEventHandler interface {
	CreateEvent(w http.ResponseWriter, r *http.Request)
}

type createEventHandler struct {
	service service.EventService
}

func NewCreateHandler(eventService service.EventService) CreateEventHandler {
	return &createEventHandler{
		service: eventService,
	}
}

func (h *createEventHandler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	data, err := PostDataFromRequest(r)
	if err != nil {
		badResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	event, err := EventFromPostData(data)
	if err != nil {
		badResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.CreateEvent(r.Context(), data.UserID, event)
	if err != nil {
		badResponse(w, http.StatusInternalServerError, err.Error())
	}

	successResponse(w, http.StatusOK)
}
