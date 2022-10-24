package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var f *os.File
	f = os.Stdin
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "STOP" {
			fmt.Println("RECEIVED STOP, NOW TERMINATING THE PROGRAM... ")
			os.Exit(0)
		}
		num, err := strconv.ParseInt(text, 0, 64)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(">", num)
	}
}
