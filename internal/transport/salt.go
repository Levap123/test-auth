package transport

import (
	"log"
	"net/http"
)

type SaltResponse struct {
	Salt string `json:"salt"`
}

func (t *Transport) generateSalt(w http.ResponseWriter, r *http.Request) {
	salt := t.service.GenerateSalt()
	if err := sendJSON(w, http.StatusOK, SaltResponse{salt}); err != nil {
		log.Printf("something whent wrong after encode: %v", err)
	}
}
