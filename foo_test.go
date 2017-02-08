package main

import (
	"testing"
)

func TestGetCDE(t *testing.T) {
	expectedResult := "CDE"

	abc := GetCDE()

	if abc != expectedResult {
		t.Errorf("Unexpected Result:", abc, "It should:", expectedResult)
	}
}