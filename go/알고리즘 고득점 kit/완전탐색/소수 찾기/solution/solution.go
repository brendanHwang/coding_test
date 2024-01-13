package solution

import (
	"math"
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

// permute 함수는 가능한 모든 숫자 조합을 생성합니다.
func permute(runes []rune, start int, result *map[int]bool) {
	if start == len(runes) {
		num, _ := strconv.Atoi(string(runes))
		(*result)[num] = true
	} else {
		for i := start; i < len(runes); i++ {
			runes[start], runes[i] = runes[i], runes[start]
			permute(runes, start+1, result)
			runes[start], runes[i] = runes[i], runes[start]
		}
	}
}

/*
주어진 숫자 문자열에서 가능한 모든 숫자 조합을 생성합니다.
각 숫자 조합에 대해 소수인지 확인합니다.
소수의 개수를 계산합니다.
*/
func solution(numbers string) int {
	runes := []rune(numbers)
	result := make(map[int]bool)
	for i := 0; i < len(runes); i++ {
		permute(runes, i, &result)
	}

	count := 0
	for num := range result {
		if isPrime(num) {
			count++
		}
	}

	return count
}
