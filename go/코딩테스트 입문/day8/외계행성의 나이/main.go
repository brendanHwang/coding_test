package main

import (
	"fmt"
	"strconv"
)

func numToAlpha(age rune) string {
	return string(age + 'a' - '0')
}

func solution(age int) string {
	ageStr := strconv.Itoa(age)
	var result string
	// fmt.Println(ageStr)
	for _, v := range ageStr {
		result += numToAlpha(v)
	}

	return result
}

func main() {
	age := 23
	result := solution(age)
	fmt.Println(result)
}
