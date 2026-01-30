package main

import "testing"

func TestHelloWorld(t *testing.T) {
	var number uint8 = 100
	_ = number

	var emoji rune = '😄'
	_ = emoji
	result := "Hello, World!"
	expected := "Hello, World!"
	if result != expected {
		t.Errorf("got %q, want %q", result, expected)
	}
}
