package transport

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (t *Transport) InitRoutes() http.Handler {
	r := chi.NewRouter()

	r.Post("/create-user", t.createUser)
	return r
}
