package main

import "fmt"

func main() {
	fmt.Println("leetcode-go")
}

// solution

// puck1609-leetcode-go
// go-util
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
