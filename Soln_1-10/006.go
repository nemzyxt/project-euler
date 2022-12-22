package main

import (
	"fmt"
	"math"
)

func sum_of_squares(n int) int {
	sum := 0
	for i:=1; i<=n; i++ {
		sum += int(math.Pow(float64(i), 2))
	}
	return sum
}

func square_of_sum(n int) int {
	sum := 0
	for i:=1; i<=n; i++ {
		sum += i
	}
	return int(math.Pow(float64(sum), 2))
}

func main() {
	n := 100
	ans := math.Abs(float64((sum_of_squares(n) - square_of_sum(n))))
	fmt.Printf("Answer : %d\n", int(ans))
}