package main

import (
	"fmt"
	"kalkulator/Calculate"
	"os"
)

func main() {

	input := os.Args

	hasil, err := Calculate.Calculate(input[1])

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("result: ", hasil[0])
}
