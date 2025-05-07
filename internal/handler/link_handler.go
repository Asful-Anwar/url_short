package handler 

import (
	"encoding/json"
	"net/http"
	"github.com/Asful-Anwar/url-shortener/internal/service"
)

type LinkHandler struct {
	Service *service.LinkService
}

func NewLinkHandler(service *service.LinkService) *LinkHandler {
	return &LinkHandler{Service: service}
}

func (h *LinkHandler) CreateShortLink(w http.ResponseWriter, r *http.Request) {
	var input struct {
	Link string `json:"link"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid Input", http.StatusBadRequest)
		return
	}
		short, err := h.Service.CreateShortLink(input.Link)
	if err != nil {
		http.Error(w, "Failed to create short link", http.StatusInternalServerError)
		return
	}
		resp := map[string]string{"short_link": short}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
}