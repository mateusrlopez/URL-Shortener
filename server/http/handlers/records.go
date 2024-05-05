package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/mateusrlopez/url-shortener-server/http/requests"
	"github.com/mateusrlopez/url-shortener-server/http/responses"
	"github.com/mateusrlopez/url-shortener-server/services"
	"github.com/rs/zerolog/log"
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

	record, err := h.service.Create(req.LongURL)
	if err != nil {
		log.Error().Err(err).Msg("error")
		http.Error(w, "failed to create record", http.StatusInternalServerError)
		return
	}

	res := responses.CreateRecord{
		ShortURLKey: record.ShortURLKey,
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
		log.Error().Err(err).Msg("error")
		http.Error(w, "failed to retrieve record by key", http.StatusInternalServerError)
		return
	}

	res := responses.FindOneRecord{
		ID:          record.ID,
		ShortURLKey: record.ShortURLKey,
		LongURL:     record.LongURL,
		Accesses:    record.Accesses,
		LastAccess:  record.LastAccess,
		CreatedAt:   record.CreatedAt,
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(&res); err != nil {
		http.Error(w, "failed to encode response into json", http.StatusInternalServerError)
		return
	}
}

func (h Records) Redirect(w http.ResponseWriter, r *http.Request) {
	key := chi.URLParam(r, "key")

	record, err := h.service.FindOneAndRegisterAccessByKey(key)
	if err != nil {
		http.Error(w, "failed to retrieve or update record", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, record.LongURL, http.StatusFound)
}
