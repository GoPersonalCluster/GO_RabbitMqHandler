package math

import "testing"

func Sum(a, b int) int {
	return a + b
}

func TestSum(t *testing.T) {
	result := Sum(2, 3)
	expected := 5

	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}
