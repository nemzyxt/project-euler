package main

import (
	"fmt"
	"math"
)

func isPrime(n int64) bool {
	var i int64 = 2
	for ; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	ans := 0
	var num int64 = 600851475143
	var x int64 = int64(math.Sqrt(float64(num)))
	for ; x > 1; x-- {
		if num%x == 0 && isPrime(x) {
			ans = int(x)
			break
		}
	}
	fmt.Printf("Answer : %d\n", ans)
}
