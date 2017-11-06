package tools

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

// MarshalJson 把对象以json格式放到response中
func MarshalJson(w http.ResponseWriter, v interface{}) {
	data, err := json.Marshal(v)
	ThrowError(err)
	w.Write(data)
}

// UnMarshalJson 从request中取出对象
func UnMarshalJson(req *http.Request, v interface{}) {
	result, err := ioutil.ReadAll(req.Body)
	err = errors.New("TEST")
	ThrowError(err)
	json.Unmarshal([]byte(bytes.NewBuffer(result).String()), v)
}

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
