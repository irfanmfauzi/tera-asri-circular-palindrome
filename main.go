package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

func isPalindrome(s string) bool {
	n := len(s)
	for i := range n / 2 {
		if s[i] != s[n-i-1] {
			return false
		}
	}
	return true
}

func isCircularPalindrome(s string) (result bool) {
	if len(s) <= 1 {
		return false
	}
	start := time.Now()
	defer func() {
		fmt.Printf("Duration :%v\n", time.Since(start))
	}()

	s = strings.ToLower(strings.ReplaceAll(s, " ", ""))

	for range len(s) {
		result = isPalindrome(s)
		if result {
			fmt.Println(result)
			return result
		}
	}
	fmt.Println(result)
	return result
}

func main() {
	fmt.Println("Enter some text:")
	var input string
	_, err := fmt.Scan(&input)
	if err != nil {
		log.Fatal(err)
		return
	}
	isCircularPalindrome(input)
}
