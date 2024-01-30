package routes

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"v2/internal/controller"
)

type Controllers struct {
	UsersController controller.Users
}

func HandlersHTTP(c Controllers) http.Handler {
	r := chi.NewRouter()
	r.HandleFunc("/users/registration", c.UsersController.RegistrationHandler)
	r.HandleFunc("/users/auth", c.UsersController.AuthHandler)
	return r
}
