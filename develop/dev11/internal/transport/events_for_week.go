package transport

import (
	"github.com/Parside01/dev11/internal/service"
	"net/http"
	"time"
)

type EventsForWeekHandler interface {
	EventsForWeek(w http.ResponseWriter, r *http.Request)
}

type eventsForWeekHandler struct {
	service service.EventService
}

func NewEventsForWeekHandler(service service.EventService) EventsForWeekHandler {
	return &eventsForWeekHandler{
		service: service,
	}
}

func (h *eventsForWeekHandler) EventsForWeek(w http.ResponseWriter, r *http.Request) {
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

	events, err := h.service.EventsByWeek(r.Context(), data.UserID, date)
	if err != nil {
		badResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	successResponseData(w, http.StatusOK, events)
}
