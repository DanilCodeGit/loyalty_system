package routes

import (
	"github.com/DanilCodeGit/loyalty_system/internal/controller"
	"github.com/go-chi/chi/v5"
	"net/http"
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
