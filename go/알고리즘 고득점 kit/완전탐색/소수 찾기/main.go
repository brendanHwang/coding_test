package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

func isPrime(number int) bool {
	if number <= 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(number))); i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

//func permute(s string) []string {
//	resultsMap := make(map[string]bool)
//	var helper func(string, string)
//	helper = func(prefix string, remaining string) {
//		if len(remaining) == 0 {
//			return
//		}
//		for i := 0; i < len(remaining); i++ {
//			current := prefix + string(remaining[i])
//			resultsMap[current] = true
//			next := remaining[:i] + remaining[i+1:]
//			helper(current, next)
//		}
//	}
//	helper("", s)
//
//	var results []string
//	for k := range resultsMap {
//		results = append(results, k)
//	}
//	sort.Strings(results) // 순서 정렬 (선택적)
//	return results
//}

func permute(s string) []string {
	resultMap := make(map[string]bool)
	var helper func(string, string)
	helper = func(prefix, remaining string) {
		if len(remaining) == 0 {
			return
		}
		for i := 0; i < len(remaining); i++ {
			current := prefix + string(remaining[i])
			resultMap[current] = true
			next := remaining[:i] + remaining[i+1:]
			helper(current, next)
		}
	}
	helper("", s)
	var result []string
	for k := range resultMap {
		result = append(result, k)
	}
	sort.Strings(result)
	return result
}

func solution(numbers string) int {
	numbersPermutations := permute(numbers)
	fmt.Printf("numbersPermutations: %v\n", numbersPermutations)

	primeCount := 0
	uniquePrimes := make(map[int]bool)
	for _, s := range numbersPermutations {
		number, _ := strconv.Atoi(s)
		if isPrime(number) && !uniquePrimes[number] {
			uniquePrimes[number] = true
			primeCount++
		}
	}

	fmt.Printf("primePermutations: %v\n", uniquePrimes)
	return primeCount
}

func main() {
	numbers := "17"
	fmt.Println(solution(numbers))
}
