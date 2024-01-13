package 구슬을_나누는_경우의_수

import (
	"testing"
)

// The function returns the correct output for valid inputs.
func TestValidInputs(t *testing.T) {
	result := solution(5, 2)
	expected := 10

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

// The function returns 1 when share is 0.
func TestShareZero(t *testing.T) {
	result := solution(5, 0)
	expected := 1

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

// The function returns 1 when share is equal to balls.
func TestShareEqualsBalls(t *testing.T) {
	result := solution(5, 5)
	expected := 1

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

// The function returns 0 when balls is 0.
func TestBallsZero(t *testing.T) {
	result := solution(0, 2)
	expected := 0

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
