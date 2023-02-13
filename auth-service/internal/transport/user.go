package transport

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/Levap123/auth-service/internal/service"
	"github.com/go-chi/chi/v5"
)

type createUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (t *Transport) createUser(w http.ResponseWriter, r *http.Request) {
	var userReq createUserRequest

	if err := readJSON(r, &userReq); err != nil {
		if err := sendJSON(w, http.StatusBadRequest, map[string]string{"message": "error in reading request"}); err != nil {
			log.Println(err)
		}
	}
	if !t.validator.IsEmailValid(userReq.Email) {
		if err := sendJSON(w, http.StatusBadRequest, map[string]string{"message": "email is invalid"}); err != nil {
			log.Println(err)
		}
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*6)
	defer cancel()
	if err := t.service.Create(ctx, userReq.Email, userReq.Password); err != nil {
		switch {
		case errors.Is(err, service.ErrUserExists):
			if err := sendJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()}); err != nil {
				log.Println(err)
			}
		default:
			if err := sendJSON(w, http.StatusInternalServerError, map[string]string{"message": err.Error()}); err != nil {
				log.Println(err)
			}
		}
		return
	}
	if err := sendJSON(w, http.StatusCreated, nil); err != nil {
		log.Println(err)
	}
}

func (t *Transport) getUserByEmail(w http.ResponseWriter, r *http.Request) {
	email := chi.URLParam(r, "email")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()
	user, err := t.service.Get(ctx, email)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrUserNotExist):
			if err := sendJSON(w, http.StatusUnauthorized, map[string]string{"message": err.Error()}); err != nil {
				log.Println(err)
			}
		default:
			if err := sendJSON(w, http.StatusInternalServerError, map[string]string{"message": err.Error()}); err != nil {
				log.Println(err)
			}
		}
		return
	}
	if err := sendJSON(w, http.StatusOK, user); err != nil {
		log.Println(err)
	}
}
