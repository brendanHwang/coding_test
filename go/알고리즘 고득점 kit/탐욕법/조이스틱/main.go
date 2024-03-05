package main

import (
	"fmt"
	"math"
)

//var alphabet = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

func distance(a1, a2 rune) int {
	dist := int(math.Abs(float64(a1 - a2)))
	if dist <= 13 {
		return dist
	} else {
		return 26 - dist
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func solution(name string) int {
	totalMoves := 0
	length := len(name)
	minMove := length - 1

	for i, char := range name {
		// 알파벳 변경을 위한 최소 조작 횟수 계산
		totalMoves += min(int(char-'A'), int('Z'-char+1))

		// 다음 문자로 넘어가기 위한 최소 커서 이동 횟수 계산
		next := i + 1
		for next < length && name[next] == 'A' {
			next++
		}

		// 현재 위치에서 다시 돌아가거나 계속 진행할 때의 최소 이동 거리 계산
		minMove = min(minMove, i+length-next+min(i, length-next))
	}

	return totalMoves + minMove
}

func main() {
	fmt.Println(solution("JEROEN")) // 예상 출력: 56
	fmt.Println(solution("JAN"))    // 예상 출력: 23
}
