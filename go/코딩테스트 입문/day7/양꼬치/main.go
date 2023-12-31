package main

import "fmt"


func solution(n int, k int) int {
    quotient := n / 10
	k = k - quotient

	if k <= 0 {
		k =0
	}

	return (n * 12000) + (k * 2000)
}

func main(){

	n := 64
	k := 6
	fmt.Println(solution(n, k))
}