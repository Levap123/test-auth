package service

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/Levap123/auth-service/internal/models"
	"github.com/Levap123/auth-service/internal/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	repo repository.UserRepo
}

func NewUser(repo repository.UserRepo) *User {
	return &User{
		repo: repo,
	}
}

func (u *User) Create(ctx context.Context, email, password string) error {
	userModel, err := u.repo.Get(ctx, email)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			return err
		}
	}
	if userModel != nil {
		return ErrUserExists
	}
	httpClient := &http.Client{}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "http://salt-gen-service:5000/generate-salt", nil)
	if err != nil {
		return fmt.Errorf("service - create user - new request -  %w", err)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("service - create user - make request - %w", err)
	}
	defer resp.Body.Close()
	var salt models.Salt
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("service - create user - read response - %w", err)
	}
	fmt.Println(string(b))
	if err := json.Unmarshal(b, &salt); err != nil {
		return fmt.Errorf("serice - create user - unmarshall salt - %w", err)
	}

	hash := md5.New()
	hash.Write([]byte(password))
	password = fmt.Sprintf("%x", hash.Sum([]byte(salt.Salt)))

	user := &models.User{
		Email:    email,
		Salt:     salt.Salt,
		Password: password,
	}

	if err := u.repo.Create(ctx, user); err != nil {
		return err
	}

	return nil
}

func (u *User) Get(ctx context.Context, email string) (*models.User, error) {
	user, err := u.repo.Get(ctx, email)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, ErrUserNotExist
		}
		return nil, err
	}
	return user, nil
}
