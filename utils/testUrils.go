package utils

import "testing"

func AssertIntEquals(t *testing.T, expected int, value int) {
	if expected == value {
		return
	}

	t.Errorf("Failed to assert %d equals expected %d", value, expected)
}

func AssertStringEquals(t *testing.T, expected string, value string) {
	if expected == value {
		return
	}

	t.Errorf("Failed to assert %s equals expected %s", value, expected)
}
