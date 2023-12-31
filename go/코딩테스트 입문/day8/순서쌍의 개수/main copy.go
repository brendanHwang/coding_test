package main

import "fmt"

func factor(n int) map[int]int {
	if n == 1 {
		return map[int]int{}
	}

	factors := map[int]int{}
	for i := 2; i*i <= n; i++ {
		for n%i == 0 {
			factors[i]++
			n /= i
		}
	}

	if n > 1 {
		factors[n]++
	}

	return factors
}

func solution(n int) int {
	ans := 1
	factorMap := factor(n)
	for _, value := range factorMap {
		ans *= (value + 1)
	}

	return ans
}

func main() {
	fmt.Println(solution(20))
}
