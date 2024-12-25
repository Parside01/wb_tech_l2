package transport

import (
	"github.com/Parside01/dev11/internal/service"
	"net/http"
)

type DeleteEventHandler interface {
	DeleteEvent(w http.ResponseWriter, r *http.Request)
}

type deleteEventHandler struct {
	service service.EventService
}

func NewDeleteEventHandler(service service.EventService) DeleteEventHandler {
	return &deleteEventHandler{
		service: service,
	}
}

func (h *deleteEventHandler) DeleteEvent(w http.ResponseWriter, r *http.Request) {
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

	if err = h.service.DeleteEvent(r.Context(), data.UserID, event); err != nil {
		badResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	successResponse(w, http.StatusOK)
}
