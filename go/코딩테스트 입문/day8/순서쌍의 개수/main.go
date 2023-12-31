package main

import "fmt"

func factor(n int) map[int]int {
    // n이 1이면 빈 맵을 반환합니다.
    if n == 1 {
        return map[int]int{}
    }

    // n의 약수를 찾습니다.
    factors := map[int]int{}
    for i := 2; i*i <= n; i++ {
        for n % i == 0 {
            factors[i]++
            n /= i
        }
    }

    // n이 소수인 경우 1을 추가합니다.
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


func main(){
	fmt.Println(solution(20))
}