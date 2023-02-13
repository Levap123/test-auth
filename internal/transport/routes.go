package transport

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (t *Transport) InitRoutes() http.Handler {
	r := chi.NewRouter()

	r.Post("/generate-salt", t.generateSalt)
	return r
}
