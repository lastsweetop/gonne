package tools

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	jsoniter "github.com/json-iterator/go"
)

func TestReverse(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
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

func BenchmarkReverse(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Reverse("s string")
	}
}

func BenchmarkReverseParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Reverse("s string")
		}
	})
}

func ExampleReverse() {
	fmt.Println(Reverse("Hello, world"))
	fmt.Println(Reverse("Hello, 世界"))
	// Unordered Output: 界世 ,olleH
	// dlrow ,olleH
}

type ColorGroup struct {
	ID     int
	Name   string
	Colors []string
}

var group = ColorGroup{
	ID:     1,
	Name:   "Reds",
	Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
}

func BenchmarkStdJson(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			b, err := json.Marshal(group)
			if err != nil {
				fmt.Println(b)
			}
		}
	})
}

func BenchmarkIterJson(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			b, err := jsoniter.Marshal(group)
			if err != nil {
				fmt.Println(b)
			}
		}
	})
}

func TestMain(m *testing.M) {
	fmt.Println("init")
	os.Exit(m.Run())
}
