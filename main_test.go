package main

import "testing"

func TestGenText(t *testing.T) {
	result := genText()
	if result != "Hello World!" {
		t.Fatalf("genText() = %v, want %v", result, "Hello World!")
	}
}
