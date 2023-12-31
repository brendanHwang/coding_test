package main

import (
	"fmt"
	"strconv"
)

func solution(intStrs []string, k int, s int, l int) []int {
	answer := make([]int, 0)
	for _, intStr := range intStrs {
		subStr := intStr[s : s+l]
		subStrInt, _ := strconv.Atoi(subStr)
		if subStrInt > k {
			answer = append(answer, subStrInt)
		}
	}
	return answer	
}

func main() {
	intStrs := []string{"0123456789", "9876543210", "9999999999999"}
	k, s, l := 50000, 5, 5
	fmt.Println(solution(intStrs, k, s, l))
}
