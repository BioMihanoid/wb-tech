package http

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"wb-tech/l2/18/internal/model"
)

type ServiceInterface interface {
	CreateEvent(userID int64, date time.Time, text string) (model.Event, error)
	UpdateEvent(userID, id int64, date time.Time, text string) (model.Event, error)
	DeleteEvent(userID, id int64) error
	EventsForDay(userID int64, date time.Time) ([]model.Event, error)
	EventsForWeek(userID int64, date time.Time) ([]model.Event, error)
	EventsForMonth(userID int64, date time.Time) ([]model.Event, error)
}

type Handler struct {
	service ServiceInterface
}

const (
	LayoutDate = "2006-01-02"
)

func NewHandler(service ServiceInterface) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	userID, date, event, ok := parseEventForm(w, r)
	if !ok {
		return
	}

	e, err := h.service.CreateEvent(userID, date, event)
	if err != nil {
		http.Error(w, `{"error":"business error"}`, http.StatusServiceUnavailable)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]any{"result": e})
	if err != nil {
		http.Error(w, `{"error":"encode json"}`, http.StatusServiceUnavailable)
		return
	}
}

func (h *Handler) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	userID, date, event, ok := parseEventForm(w, r)
	if !ok {
		return
	}

	id, err := strconv.ParseInt(r.Form.Get("user_id"), 10, 64)
	if err != nil {
		http.Error(w, `{"error":"invalid id"}`, http.StatusBadRequest)
		return
	}

	e, err := h.service.UpdateEvent(userID, id, date, event)
	if err != nil {
		http.Error(w, `{"error":"business error"}`, http.StatusServiceUnavailable)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]any{"result": e})
	if err != nil {
		http.Error(w, `{"error":"encode json"}`, http.StatusServiceUnavailable)
		return
	}
}

func (h *Handler) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, `{"error":"bad request"}`, http.StatusBadRequest)
		return
	}

	userID, err := strconv.ParseInt(r.Form.Get("user_id"), 10, 64)
	if err != nil {
		http.Error(w, `{"error":"invalid user_id"}`, http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(r.Form.Get("user_id"), 10, 64)
	if err != nil {
		http.Error(w, `{"error":"invalid id"}`, http.StatusBadRequest)
		return
	}

	err = h.service.DeleteEvent(userID, id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]string{"result": "deleted"})
	if err != nil {
		http.Error(w, `{"error":"encode json"}`, http.StatusServiceUnavailable)
	}
}

func (h *Handler) EventsForDay(w http.ResponseWriter, r *http.Request) {
	userID, date, ok := parseQuery(w, r)
	if !ok {
		return
	}

	events, err := h.service.EventsForDay(userID, date)
	if err != nil {
		http.Error(w, `{"error":"internal error"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(map[string]any{"result": events})
	w.WriteHeader(http.StatusOK)
	if err != nil {
		http.Error(w, `{"error":"encode json"}`, http.StatusServiceUnavailable)
		return
	}
}

func (h *Handler) EventsForWeek(w http.ResponseWriter, r *http.Request) {
	userID, date, ok := parseQuery(w, r)
	if !ok {
		return
	}

	events, err := h.service.EventsForWeek(userID, date)
	if err != nil {
		http.Error(w, `{"error":"internal error"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]any{"result": events})
	if err != nil {
		http.Error(w, `{"error":"encode json"}`, http.StatusServiceUnavailable)
		return
	}
}

func (h *Handler) EventsForMonth(w http.ResponseWriter, r *http.Request) {
	userID, date, ok := parseQuery(w, r)
	if !ok {
		return
	}

	events, err := h.service.EventsForMonth(userID, date)
	if err != nil {
		http.Error(w, `{"error":"internal error"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]any{"result": events})
	if err != nil {
		http.Error(w, `{"error":"encode json"}`, http.StatusServiceUnavailable)
		return
	}
}

func parseEventForm(w http.ResponseWriter, r *http.Request) (int64, time.Time, string, bool) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, `{"error":"bad request"}`, http.StatusBadRequest)
		return 0, time.Time{}, "", false
	}

	userID, err := strconv.ParseInt(r.Form.Get("user_id"), 10, 64)
	if err != nil {
		http.Error(w, `{"error":"invalid user_id"}`, http.StatusBadRequest)
		return 0, time.Time{}, "", false
	}

	date, err := time.Parse(LayoutDate, r.Form.Get("date"))
	if err != nil {
		http.Error(w, `{"error":"invalid date"}`, http.StatusBadRequest)
		return 0, time.Time{}, "", false
	}

	event := r.Form.Get("event")
	return userID, date, event, true
}

func parseQuery(w http.ResponseWriter, r *http.Request) (int64, time.Time, bool) {
	userIDStr := r.URL.Query().Get("user_id")
	dateStr := r.URL.Query().Get("date")

	if userIDStr == "" || dateStr == "" {
		http.Error(w, `{"error":"user_id and date are required"}`, http.StatusBadRequest)
		return 0, time.Time{}, false
	}

	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		http.Error(w, `{"error":"invalid user_id"}`, http.StatusBadRequest)
		return 0, time.Time{}, false
	}

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		http.Error(w, `{"error":"invalid date"}`, http.StatusBadRequest)
		return 0, time.Time{}, false
	}

	return userID, date, true
}
