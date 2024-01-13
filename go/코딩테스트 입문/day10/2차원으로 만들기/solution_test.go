package _차원으로_만들기

import (
	"reflect"
	"testing"
)

func Test1(t *testing.T) {
	num_list := []int{1, 2, 3, 4, 5, 6, 7, 8}
	n := 2
	result := solution(num_list, n)
	expected := [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}

	if !reflect.DeepEqual(result, expected) { //Go의 표준 테스팅 패키지는 reflect.DeepEqual 함수를 사용하여 두 슬라이스 또는 배열이 동일한지 여부를 확인할 수 있습니다.
		t.Errorf("solution(%v, %d) = %v; want %v", num_list, n, result, expected)
	}
}

func Test2(t *testing.T) {
	num_list := []int{100, 95, 2, 4, 5, 6, 18, 33, 948}
	n := 3
	result := solution(num_list, n)
	expected := [][]int{{100, 95, 2}, {4, 5, 6}, {18, 33, 948}}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("solution(%v, %d) = %v; want %v", num_list, n, result, expected)
	}
}
