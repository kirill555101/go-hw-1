package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Invalid arguments number")
		return
	}

	postfixNotation, err := GetPostfixNotation(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}

	result, err := Calculate(postfixNotation)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}
