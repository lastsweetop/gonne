package tools

import (
	"fmt"
	"net/http"

	"github.com/lastsweetop/gonne/model"
	"github.com/pkg/errors"
)

func ThrowError(err error) {
	if err != nil {
		panic(errors.WithStack(err))
	}
}

func CatchError(w http.ResponseWriter, req *http.Request) {
	if r := recover(); r != nil {
		resp := &model.Resp{Code: "9001", Msg: r.(error).Error()}
		MarshalJson(w, resp)
		fmt.Printf("%+v\n", r.(error))
	}
}
