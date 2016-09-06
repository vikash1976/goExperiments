package exporter

import (
	"fmt"
	"testing"
)

func TestGreeter(t *testing.T) {
	cases := []struct {
		in1, in2, want string
	}{
		{"es", "Vikash", "Hola Vikash"},
		{"en", "Gudia", "Hello Gudia"},
		{"", "Aryan", "Hi Aryan"},
	}
	for _, c := range cases {
		got := Greeter(c.in1, c.in2)
		if got != c.want {
			t.Errorf("get(%q and %q) == %q, want %q", c.in1, c.in2, got, c.want)
		} else {
			fmt.Printf("Test successful! - get(%q and %q) == %q, want %q\n", c.in1, c.in2, got, c.want)
		}
	}

}
