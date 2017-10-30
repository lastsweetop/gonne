package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewAPIMux() *mux.Router {
	r := mux.NewRouter()
	s := r.PathPrefix("/api").Subrouter()
	initUserApi(s)

	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("web/"))))
	return r
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello mux"))
}
