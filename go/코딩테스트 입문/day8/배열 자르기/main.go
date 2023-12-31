package main

import "fmt"

func solution(numbers []int, num1 int, num2 int) []int {
	return numbers[num1 : num2+1]
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	num1 := 1
	num2 := 3
	result := solution(numbers, num1, num2)

	fmt.Println(result)
}
