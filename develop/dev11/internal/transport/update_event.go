package transport

import (
	"github.com/Parside01/dev11/internal/service"
	"net/http"
)

type UpdateEventHandler interface {
	UpdateEvent(w http.ResponseWriter, r *http.Request)
}

type updateEventHandler struct {
	service service.EventService
}

func NewUpdateHandler(eventService service.EventService) UpdateEventHandler {
	return &updateEventHandler{
		service: eventService,
	}
}

func (h *updateEventHandler) UpdateEvent(w http.ResponseWriter, r *http.Request) {
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

	err = h.service.UpdateEvent(r.Context(), data.UserID, event)
	if err != nil {
		badResponse(w, http.StatusInternalServerError, err.Error())
	}

	successResponse(w, http.StatusOK)
}
