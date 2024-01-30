package controller

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"net/http"
	"strings"
	"v2/internal/auth"
	"v2/internal/domain"
)

type UserService interface {
	Registration(ctx context.Context, user domain.Users) error
	Authentication(ctx context.Context, user domain.Users) error
}

type Users struct {
	s UserService
}

func NewUsers(s UserService) Users {
	return Users{s: s}
}

func (u Users) Registration(ctx context.Context, user domain.Users) (string, error) {
	err := u.s.Registration(ctx, user)
	if err != nil {
		return "", err
	}
	token, err := auth.GenerateJWT(user)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (u Users) RegistrationHandler(w http.ResponseWriter, r *http.Request) {
	var newUser domain.Users
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if newUser.Login == "" || newUser.Password == "" {
		http.Error(w, "Login and password are required", http.StatusBadRequest)
		return
	}

	token, err := u.Registration(context.TODO(), newUser)
	if err != nil {
		switch {
		case strings.Contains(err.Error(), "user already exists"):
			http.Error(w, "Login already exists", http.StatusConflict)
		default:
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	err = auth.ValidateToken(w, r)
	if err != nil {
		return
	}
	w.Write([]byte(token))
}

func (u Users) Authentication(ctx context.Context, user domain.Users) error {
	err := u.s.Authentication(ctx, user)
	if err != nil {
		return errors.WithMessage(err, "controller")
	}
	return nil
}

func (u Users) AuthHandler(w http.ResponseWriter, r *http.Request) {
	var newUser domain.Users
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if newUser.Login == "" || newUser.Password == "" {
		http.Error(w, "Login and password are required", http.StatusBadRequest)
		return
	}

	err = u.s.Authentication(context.TODO(), newUser)
	if err != nil {
		http.Error(w, "auth user", http.StatusUnauthorized)
		return
	}

	err = auth.ValidateToken(w, r)
	if err != nil {
		http.Error(w, "access is denied", http.StatusUnauthorized)
		return
	}
	w.Write([]byte("User successfully authenticated"))
}
