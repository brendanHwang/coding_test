package main

import (
	"testing"
)

func Test1(t *testing.T) {
	expect := 5
	result := Solution(5, []int{2, 4}, []int{1, 3, 5})
	if expect != result {
		t.Errorf("expected %v but %v\n", expect, result)
	}
}

func Test2(t *testing.T) {
	expect := 4
	result := Solution(5, []int{2, 4}, []int{3})
	if expect != result {
		t.Errorf("expected %v but %v\n", expect, result)
	}
}

func Test3(t *testing.T) {
	expect := 2
	result := Solution(3, []int{3}, []int{1})
	if expect != result {
		t.Errorf("expected %v but %v\n", expect, result)
	}
}

func Test4(t *testing.T) {
	expect := 2
	result := Solution(3, []int{3, 4}, []int{3})
	if expect != result {
		t.Errorf("expected %v but %v\n", expect, result)
	}
}

func Test5(t *testing.T) {
	expect := 2
	result := Solution(3, []int{2, 3}, []int{3})
	if expect != result {
		t.Errorf("expected %v but %v\n", expect, result)
	}
}

func Test6(t *testing.T) {
	expect := 3
	result := Solution(3, []int{2}, []int{2, 3})
	if expect != result {
		t.Errorf("expected %v but %v\n", expect, result)
	}
}

func Test7(t *testing.T) {
	expect := 3
	result := Solution(3, []int{2}, []int{2, 1})
	if expect != result {
		t.Errorf("expected %v but %v\n", expect, result)
	}
}

func Test8(t *testing.T) {
	expect := 3
	result := Solution(4, []int{3, 4}, []int{4})
	if expect != result {
		t.Errorf("expected %v but %v\n", expect, result)
	}
}

func Test9(t *testing.T) {
	expect := 3
	result := Solution(4, []int{1, 2, 3, 4}, []int{1, 2, 3})
	if expect != result {
		t.Errorf("expected %v but %v\n", expect, result)
	}
}

func Test10(t *testing.T) {
	expect := 4
	result := Solution(4, []int{1, 2, 4}, []int{1, 2, 3, 4})
	if expect != result {
		t.Errorf("expected %v but %v\n", expect, result)
	}
}

func Test11(t *testing.T) {
	expect := 4
	result := Solution(4, []int{1, 4}, []int{1, 2, 3, 4})
	if expect != result {
		t.Errorf("expected %v but %v\n", expect, result)
	}
}

func Test13(t *testing.T) {
	expect := 4
	result := Solution(4, []int{2, 4}, []int{1, 2, 3, 4})
	if expect != result {
		t.Errorf("expected %v but %v\n", expect, result)
	}
}

func Test14(t *testing.T) {
	expect := 4
	result := Solution(4, []int{2, 3, 4}, []int{1, 2, 3, 4})
	if expect != result {
		t.Errorf("expected %v but %v\n", expect, result)
	}
}

func Test15(t *testing.T) {
	expect := 2
	result := Solution(4, []int{2, 3, 4}, []int{1})
	if expect != result {
		t.Errorf("expected %v but %v\n", expect, result)
	}
}

func Test16(t *testing.T) {
	expect := 1
	result := Solution(4, []int{2, 3, 4}, []int{0})
	if expect != result {
		t.Errorf("expected %v but %v\n", expect, result)
	}
}

func Test17(t *testing.T) {
	expect := 1
	result := Solution(4, []int{2, 3, 4}, []int{6})
	if expect != result {
		t.Errorf("expected %v but %v\n", expect, result)
	}
}
func Test18(t *testing.T) {
	expect := 4
	result := Solution(6, []int{2, 3, 4, 5}, []int{1, 6})
	if expect != result {
		t.Errorf("expected %v but %v\n", expect, result)
	}
}
