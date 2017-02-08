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

func TestGetDEF(t *testing.T) {
	expectedResult := "DEF"

	abc := GetDEF()

	if abc != expectedResult {
		t.Errorf("Unexpected Result:", abc, "It should:", expectedResult)
	}
}

func TestGetEFG(t *testing.T) {
    expectedResult := "EFG"

    abc := GetEFG()

    if abc != expectedResult {
        t.Errorf("Unexpected Result:", abc, "It should:", expectedResult)
    }
}
