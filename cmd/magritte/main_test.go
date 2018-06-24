package main

import (
	"testing"
)

func TestGetTextSpeaksFrench(t *testing.T) {
	expected := "La Trahison des Images"
	actual := getText()
	if actual != expected {
		t.Errorf("getText returned unexpected body: got %v want %v", actual,
			expected)
	}
}
