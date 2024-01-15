package routes

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Controllers struct {
}

func HandlerHTTP(c Controllers) http.Handler {
	muxerHTTP := mux.NewRouter()
	muxerHTTP.Handle("/users/registration")

	return muxerHTTP
}
