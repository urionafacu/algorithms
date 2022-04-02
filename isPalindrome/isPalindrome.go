package main

import "fmt"

var print = fmt.Println

func isPalindrome(s string) bool {
	var mirror string
	for i := len(s) - 1; i >= 0; i-- {
		mirror = mirror + string(s[i])
	}

	print(mirror, s)

	return s == mirror
}

func main() {
	const ala string = "abca"
	print(isPalindrome(ala))
}
