package main 

import (
	"fmt"
)

func isPrime(n int) bool {
	for i:=2; i<n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	ans := 0
	counter := 0
	for i:=2; ;i++ {
		if isPrime(i) {
			counter++
			if counter == 10001 {
				ans = i
				break
			}
		}
	}
	fmt.Printf("Answer : %d\n", ans)
}