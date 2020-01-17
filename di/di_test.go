package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Niels")

	got := buffer.String()
	want := "Hello, Niels"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
