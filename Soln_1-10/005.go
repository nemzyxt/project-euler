package main 

import (
	"fmt"
)

func isEvenlyDivisible(n int) bool {
	for i:=1; i<=20; i++ {
		if n%i != 0 {
			return false
		}
	}
	return true
}

func main() {
	ans := 0
	for i:=21; ;i++ {
		if isEvenlyDivisible(i) {
			ans = i
			break
		}
	}
	fmt.Printf("Answer : %d\n", ans)
}