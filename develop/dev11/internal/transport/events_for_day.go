package transport

import (
	"github.com/Parside01/dev11/internal/service"
	"net/http"
	"time"
)

type EventsForDayHandler interface {
	EventsForDay(w http.ResponseWriter, r *http.Request)
}

type eventsForDayHandler struct {
	service service.EventService
}

func NewEventsForDayHandler(service service.EventService) EventsForDayHandler {
	return &eventsForDayHandler{
		service: service,
	}
}

func (h *eventsForDayHandler) EventsForDay(w http.ResponseWriter, r *http.Request) {
	data, err := GetDataFromRequest(r)
	if err != nil {
		badResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	date, err := time.Parse(time.RFC3339, data.Date)
	if err != nil {
		badResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	events, err := h.service.EventsByDay(r.Context(), data.UserID, date)
	if err != nil {
		badResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	successResponseData(w, http.StatusOK, events)
}
