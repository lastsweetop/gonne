package tools

import (
	"net/http"
	"testing"
)

func TestReverse(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}
	for _, c := range cases {
		got := Reverse(c.in)
		if got != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestMarshalJson(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		v interface{}
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MarshalJson(tt.args.w, tt.args.v)
		})
	}
}

func TestUnMarshalJson(t *testing.T) {
	type args struct {
		req *http.Request
		v   interface{}
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UnMarshalJson(tt.args.req, tt.args.v)
		})
	}
}
