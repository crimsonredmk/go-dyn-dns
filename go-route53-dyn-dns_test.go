package main

import "testing"

func TestFour(t *testing.T) {
	four := 2 + 2
	if four != 4 {
		t.Error("Expected 4, got ", four)
	}
}
