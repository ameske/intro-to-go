package main

import "testing"

func TestAdd(t *testing.T) {
	result := Add(2, 2)

	if result != 4 {
		t.Error("Expected 4 but got", result)
	}

	result = Add(0, 4)

	if result != 4 {
		t.Error("Expected 4 but got", result)
	}
}
