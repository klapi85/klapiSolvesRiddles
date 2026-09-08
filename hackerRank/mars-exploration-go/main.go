package main

import (
	"fmt"
)

/*
 * W. Kuchta 2026
 * https://www.hackerrank.com/challenges/mars-exploration
 *
 */

func marsExploration(s string) int {
	result := 0
	for i := 0; i < len(s); i += 3 {
		if "S" != s[i:i+1] {
			result++
		}
		if "O" != s[i+1:i+2] {
			result++
		}
		if "S" != s[i+2:i+3] {
			result++
		}
	}
	return result
}

func main() {
	s := "SOS456SOS"

	result := marsExploration(s)

	fmt.Println(result)
}
