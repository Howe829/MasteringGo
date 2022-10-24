package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Please give one or more float.")
		os.Exit(1)
	}
	arguments := os.Args
	min, _ := strconv.ParseFloat(arguments[1], 64)
	max, _ := strconv.ParseFloat(arguments[1], 64)

	for i := 2; i < len(arguments); i++ {
		num, _ := strconv.ParseFloat(arguments[i], 64)
		if num < min {
			min = num
		}

		if num > max {
			max = num
		}
	}

	fmt.Println("Max:", max)
	fmt.Println("Min:", min)
}
