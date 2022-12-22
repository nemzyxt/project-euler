package main

import (
	"fmt"
)

func reverseNumber(n int) int {
	rev_num := 0
	for n > 0 {
		rev_num = rev_num*10 + n%10
		n /= 10
	}
	return rev_num
}

func isPalindrome(n int) bool {
	if reverseNumber(n) == n {
		return true
	}
	return false
}

func main() {
	ans := 0
	for i := 999; i >= 100; i-- {
		for j := 999; j >= 100; j-- {
			tmp := i * j
			if isPalindrome(tmp) && i != j {
				fmt.Println(i, " ", j)
				ans = tmp
				break
			}
		}
	}
	fmt.Printf("Answer : %d\n", ans)
}
