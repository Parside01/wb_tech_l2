package transport

import (
	"github.com/Parside01/dev11/internal/service"
	"net/http"
	"time"
)

type EventsForMonthHandler interface {
	EventsForMonth(w http.ResponseWriter, r *http.Request)
}

type eventsForMonthHandler struct {
	service service.EventService
}

func NewEventsForMonthHandler(service service.EventService) EventsForMonthHandler {
	return &eventsForMonthHandler{
		service: service,
	}
}

func (h *eventsForMonthHandler) EventsForMonth(w http.ResponseWriter, r *http.Request) {
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
