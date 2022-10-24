package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Please give one or more argument!")
		os.Exit(-1)
	}
	sum := 0.0
	argument := os.Args
	for _, numStr := range argument[1:] {
		num, err := strconv.ParseFloat(numStr, 64)
		if err != nil {
			log.Panicln(err)

		}
		sum = sum + num
	}
	fmt.Println("Sum: ", sum)
	fmt.Println("Average:", sum/float64(len(argument[1:])))
}
