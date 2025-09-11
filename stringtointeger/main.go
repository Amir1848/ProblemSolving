package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println(myAtoi("42"))
	fmt.Println(myAtoi(" -042"))
	fmt.Println(myAtoi("+11337c0d3"))
	fmt.Println(myAtoi("0-1"))
	fmt.Println(myAtoi("00000-42a1234"))
}

func myAtoi(s string) int {
	s = strings.TrimSpace(s)

	const zeroCode = int('0')
	const nineCode = int('9')

	var numbers []int
	isPositive := true
	isNumberStarted := false

	for _, c := range s {
		charCode := int(c)
		if c == '-' && !isNumberStarted {
			isNumberStarted = true
			isPositive = false
		} else if c == '+' && !isNumberStarted {
			isNumberStarted = true
		} else if charCode < zeroCode || charCode > nineCode {
			break
		} else {
			isNumberStarted = true
			convertedDigit := int(c) - zeroCode
			numbers = append(numbers, convertedDigit)
		}

	}

	result := 0
	pow := 0

	for i := len(numbers) - 1; i >= 0; i-- {
		if numbers[i] != 0 {
			

			paye := int64(math.Pow10(pow)) * int64(numbers[i])
	
			result += paye
		} 

		pow++
	}

	if !isPositive {
		result *= -1
	}

	return result
}
