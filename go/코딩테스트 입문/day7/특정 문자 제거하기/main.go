package main

import "fmt"


func solution(s string, letter string) string {

	var result string
	for _, v := range s {
		if string(v) != letter {
			result += string(v)
		}
	}
	return result
}
func main(){

	s := "baabaa"
	letter := "b"
	
	// letter를 제거한 s를 반환

	fmt.Println(solution(s, letter))
}