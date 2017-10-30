package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func initUserApi(r *mux.Router) {
	s := r.PathPrefix("/user").Subrouter()
	s.HandleFunc("/list", UserListHandler)
	s.HandleFunc("/add", UserAddHandler)
}

func UserListHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("user list"))
}

func UserAddHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("user add"))
}
