package main

import (
	"fmt"
	"math"
)

//Given a signed 32-bit integer x, return x with its digits reversed.
//If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.
//site has problem.Looks like the test-values have invalid test cases for x:
// 1534236469, 1563847412 and -1563847412
// Because all these numbers are within the required range [-2^31 ; 2^31 -1]

func main() {
	fmt.Print(reverse(-547))
}

func reverse(x int) int {
	digits := []int{}

	for x != 0 {
		n := x % 10
		x = x / 10
		digits = append(digits, n)
	}
	var reversedNumber int

	var tavan = int(math.Pow(10, float64(len(digits)-1)))
	for _, val := range digits {
		reversedNumber = reversedNumber + (val * tavan)
		tavan /= 10
	}

	return reversedNumber
}
