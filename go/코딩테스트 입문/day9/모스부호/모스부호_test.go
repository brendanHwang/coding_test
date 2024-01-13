package 모스부호

import "testing"

func Test1(t *testing.T) {
	result := Solution(".... . .-.. .-.. ---")
	expected := "hello"

	if result != expected {
		t.Errorf("Solution() = %s, want %s", result, expected)
	}
}

func Test2(t *testing.T) {
	result := Solution(".--. -.-- - .... --- -.")
	expected := "python"

	if result != expected {
		t.Errorf("Solution() = %s, want %s", result, expected)
	}
}
