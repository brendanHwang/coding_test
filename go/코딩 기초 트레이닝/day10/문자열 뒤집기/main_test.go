package 문자열_뒤집기

import "testing"

func SolutionTest1(t *testing.T) {
	expect := "ProgrammerS123"
	result := Solution(
		"Progra21Sremm3",
		6,
		12,
	)

	if result != expect {
		t.Errorf("Expect %v, but %v", expect, result)
	}
}

func SolutionTest2(t *testing.T) {
	expect := "Stanley1yelnatS"
	result := Solution(
		"Stanley1yelnatS",
		4,
		10,
	)

	if result != expect {
		t.Errorf("Expect %v, but %v", expect, result)
	}
}
