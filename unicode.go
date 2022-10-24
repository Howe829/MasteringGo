package main

import (
	"fmt"
	"unicode"
)

func main() {
	const sL = "\x68\x6f\x77\x65"

	for i := 0; i < len(sL); i++ {
		if unicode.IsPrint(rune(sL[i])) {
			fmt.Printf("%c", sL[i])
		} else {
			fmt.Println("Not Printable!")
		}
	}

	fmt.Println()
}
