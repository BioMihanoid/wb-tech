package http

import "net/http"

func NewMux(handler *Handler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/create_event", handler.CreateEvent)
	mux.HandleFunc("/update_event", handler.UpdateEvent)
	mux.HandleFunc("/delete_event", handler.DeleteEvent)
	mux.HandleFunc("/events_for_day", handler.EventsForDay)
	mux.HandleFunc("/events_for_week", handler.EventsForWeek)
	mux.HandleFunc("/events_for_month", handler.EventsForMonth)

	return mux
}
