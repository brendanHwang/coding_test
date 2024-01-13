package 모스부호

import "strings"

// Solution translates a string of Morse code into English letters.
//
// letter: the input string of Morse code to be translated.
// Returns the corresponding English letters.
func Solution(letter string) string {
	morse := map[string]string{
		".-":   "a",
		"-...": "b",
		"-.-.": "c",
		"-..":  "d",
		".":    "e",
		"..-.": "f",
		"--.":  "g",
		"....": "h",
		"..":   "i",
		".---": "j",
		"-.-":  "k",
		".-..": "l",
		"--":   "m",
		"-.":   "n",
		"---":  "o",
		".--.": "p",
		"--.-": "q",
		".-.":  "r",
		"...":  "s",
		"-":    "t",
		"..-":  "u",
		"...-": "v",
		".--":  "w",
		"-..-": "x",
		"-.--": "y",
		"--..": "z",
	}
	// string을 공백을 기준으로 split
	strArr := strings.Split(letter, " ")
	ans := ""

	for _, v := range strArr {
		ans += morse[v]
	}

	return ans
}
