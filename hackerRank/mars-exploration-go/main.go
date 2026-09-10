package main

import (
	"fmt"
	"os"
)

/*
 * W. Kuchta 2026
 * https://www.hackerrank.com/challenges/mars-exploration
 *
 */

func checkLetter(so string, letter string) int {
	result := 0

	if so != letter {
		result++
	}
	return result
}

func marsExploration(s string) int {
	result := 0
	for i := 0; i < len(s); i += 3 {
		result += checkLetter("S", s[i:i+1])
		result += checkLetter("O", s[i+1:i+2])
		result += checkLetter("S", s[i+2:i+3])
	}
	return result
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please add an argument like 'SOSSOS'!")
		return
	}

	s := os.Args[1]

	result := marsExploration(s)

	fmt.Println(result)
}
