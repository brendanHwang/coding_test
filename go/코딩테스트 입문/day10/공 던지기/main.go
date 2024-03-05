package main

import "fmt"

func solution(numbers []int, k int) int {
	cnt := 0
	for i := 0; ; i += 2 {
		i %= len(numbers)

		cnt++
		if cnt == k {
			return numbers[i]
		}
	}
}

func main() {
	fmt.Println(solution([]int{1, 2, 3, 4}, 2))
	fmt.Println(solution([]int{1, 2, 3, 4, 5, 6}, 5))
	fmt.Println(solution([]int{1, 2, 3}, 3))

}
