package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	inputString, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	startPosition := 0
	result := []rune{}
	if inputString[0] == '-' {
		startPosition = 1
		result = append(result, '-')
	}

	for i := len([]rune(inputString)) - 3; i >= startPosition; i-- {
		result = append(result, []rune(inputString)[i])
	}

	fmt.Println(string(result))
}
