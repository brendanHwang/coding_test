package 개미군단

import "testing"

func Test1(t *testing.T) {

	result := Solution(23)
	expected := 5

	if result != expected {
		t.Errorf("Solution() = %d, want %d", result, expected)
	}
}

func Test2(t *testing.T) {

	result := Solution(24)
	expected := 6

	if result != expected {
		t.Errorf("Solution() = %d, want %d", result, expected)
	}
}

func Test3(t *testing.T) {

	result := Solution(999)
	expected := 201

	if result != expected {
		t.Errorf("Solution() = %d, want %d", result, expected)
	}
}
