package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lastsweetop/gonne/model"
	"github.com/lastsweetop/gonne/tools"
)

func initUserApi(r *mux.Router) {
	s := r.PathPrefix("/user").Subrouter()
	s.HandleFunc("/list", UserListHandler)
	s.HandleFunc("/add", UserAddHandler)
	s.HandleFunc("/login", UserLoginndler)
}

func UserLoginndler(w http.ResponseWriter, r *http.Request) {
	user := &model.User{}
	tools.UnMarshalJson(r, user)
	resp := &model.Resp{Code: "1001", Msg: "账号或者密码错误"}
	if user.Username == "sweetop" && user.Password == "123456" {
		resp = &model.Resp{Code: "0", Msg: "success"}
	}
	tools.MarshalJson(w, resp)
}

func UserListHandler(w http.ResponseWriter, r *http.Request) {
	resp := &model.Resp{Code: "0", Msg: "success"}
	users := make([]model.User, 0)
	resp.Data = append(users, model.User{Username: "sweetop"})
	tools.MarshalJson(w, resp)
}

func UserAddHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("user add"))
}
