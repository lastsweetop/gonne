package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lastsweetop/gonne/tools"
)

type MyRouter struct {
	*mux.Router
}

func (r *MyRouter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	defer tools.CatchError(w, req)
	r.Router.ServeHTTP(w, req)
}

func NewAPIMux() http.Handler {
	r := &MyRouter{mux.NewRouter()}
	s := r.PathPrefix("/api").Subrouter()
	initUserApi(s)

	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("web/"))))
	return r
}
