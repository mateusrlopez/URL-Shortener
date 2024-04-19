package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/mateusrlopez/url-shortener-server/http/requests"
	"github.com/mateusrlopez/url-shortener-server/http/responses"
	"github.com/mateusrlopez/url-shortener-server/services"
	"net/http"
)

type Records struct {
	service *services.Records
}

func NewRecords(service *services.Records) *Records {
	return &Records{
		service: service,
	}
}

func (h Records) Create(w http.ResponseWriter, r *http.Request) {
	var req requests.CreateRecord

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "failed to decode json body", http.StatusUnprocessableEntity)
		return
	}
	defer r.Body.Close()

	if err := req.Validate(); err != nil {
		http.Error(w, "failed to validate request content", http.StatusBadRequest)
		return
	}

	key, err := h.service.Create(req.URL)
	if err != nil {
		http.Error(w, "failed to create short url record", http.StatusInternalServerError)
		return
	}

	res := responses.CreateRecord{
		Key: key,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err = json.NewEncoder(w).Encode(&res); err != nil {
		http.Error(w, "failed to encode response into json", http.StatusInternalServerError)
		return
	}
}

func (h Records) FindOneByKey(w http.ResponseWriter, r *http.Request) {
	key := chi.URLParam(r, "key")

	record, err := h.service.FindOneByKey(key)
	if err != nil {
		http.Error(w, "failed to retrieve record by key", http.StatusInternalServerError)
		return
	}

	res := responses.FindOneRecord{
		Key:        record.Key,
		URL:        record.URL,
		Accesses:   record.Accesses,
		LastAccess: record.LastAccess,
		CreatedAt:  record.CreatedAt,
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(&res); err != nil {
		http.Error(w, "failed to encode response into json", http.StatusInternalServerError)
		return
	}
}

func (h Records) Redirect(w http.ResponseWriter, r *http.Request) {
	key := chi.URLParam(r, "key")

	url, err := h.service.FindOneAndRegisterAccessByKey(key)
	if err != nil {
		http.Error(w, "failed to retrieve or update record", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, url, http.StatusFound)
}
