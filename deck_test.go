package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 28 {
		t.Errorf("Expected 28 but got %v", len(d))
	}
}
