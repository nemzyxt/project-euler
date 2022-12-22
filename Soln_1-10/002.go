package main

import (
	"fmt"
)

func main() {
	prev := 1
	curr := 2
	sum := 0
	for curr <= 4e6{
		if curr%2 == 0 {
			sum += curr
		}
		prev, curr = curr, prev+curr
	}
	fmt.Printf("Answer : %d\n", sum)
}