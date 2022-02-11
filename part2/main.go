package main

import (
	"fmt"
	"os"
	"./calc"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Invalid arguments number")
		return
	}

	postfixNotation, err := calc.GetPostfixNotation(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}

	result, err := calc.Calculate(postfixNotation)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}
