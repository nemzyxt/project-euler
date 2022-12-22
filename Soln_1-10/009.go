package main

import (
	"fmt"
	"math"
)

func isPythagoreanTriplet(a, b, c int) bool {
	x, y, z := float64(a), float64(b), float64(c)
	return (math.Pow(x,2) + math.Pow(y,2)) == math.Pow(z,2)
}

func main() {
	ans := 0
	for a:=1; a<1000; a++ {
		for b:=(a+1); b<1000; b++ {
			c := 1000-(a+b)
			if isPythagoreanTriplet(a, b, c) {
				ans = a * b * c
			}
		}
	}
	fmt.Printf("Answer : %d\n", ans)
}