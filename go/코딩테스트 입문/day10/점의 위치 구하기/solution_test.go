package 점의_위치_구하기

import "testing"

func Test1(t *testing.T) {
	dot := []int{1, 4}
	result := Solution(dot)
	expected := 1

	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}

func Test2(t *testing.T) {
	dot := []int{-7, 9}
	result := Solution(dot)
	expected := 2

	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}
