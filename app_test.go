package main

import (
	"testing"
)

func TestAddSquare(t *testing.T) {
	a := 2
	b := 3
	c := AddSquare(a, b)

	expectedResult := 13
	
	if c != expectedResult {
		t.Errorf("Unexpected Result:", c, "It should:", expectedResult)
	}
}

func TestGetABC(t *testing.T) {
	expectedResult := "ABC"

	abc := GetABC()

	if abc != expectedResult {
		t.Errorf("Unexpected Result:", abc, "It should:", expectedResult)
	}
}