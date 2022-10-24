package main

import "fmt"

const (
	Mon = iota + 1
	Tue
	Wen
	Thu
	Fri
	Sat
	Sun
)

func main() {
	fmt.Println(Mon, Tue, Wen, Thu, Fri, Sat, Sun)
}
