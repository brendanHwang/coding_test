package main

import (
	"fmt"
	"math"
)

func sumSlice(numbers []int, i int) int {
	sum := 0
	for _, v := range numbers {
		remainder := i % 2
		i = i / 2

		if remainder == 0 {
			sum += v
		} else {
			sum -= v
		}
	}

	return sum
}
func solution(numbers []int, target int) int {
	result := 0
	nPow := math.Pow(2, float64(len(numbers)))
	for i := 0; i < int(nPow); i++ {
		if target == sumSlice(numbers, i) {
			result++
		}
	}

	return result
}

func main() {
	fmt.Println(solution([]int{1, 1, 1, 1, 1}, 3))
	fmt.Println(solution([]int{4, 1, 2, 1}, 4))

}
