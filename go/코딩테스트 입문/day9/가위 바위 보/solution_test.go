package 가위_바위_보

import "testing"

func Test1(t *testing.T) {
	result := Solution("2")
	expected := "0"
	if result != expected {
		t.Errorf("Solution() = %s, want %s", result, expected)
	}
}

func Test2(t *testing.T) {
	result := Solution("205")
	expected := "052"
	if result != expected {
		t.Errorf("Solution() = %s, want %s", result, expected)
	}
}
