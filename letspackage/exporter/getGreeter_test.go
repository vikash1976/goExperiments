package exporter

import (
	"fmt"
	"testing"
)

func TestGetGreeter(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"es", "Hola"},
		{"en", "Hello"},
		{"", "Hi"},
	}
	for _, c := range cases {
		got := getGreeter(c.in)
		if got != c.want {
			t.Errorf("get(%q) == %q, want %q", c.in, got, c.want)
		} else {
			fmt.Printf("Test successful! - get(%q) == %q, want %q\n", c.in, got, c.want)
		}
	}
}
