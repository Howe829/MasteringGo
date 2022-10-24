package main

import "fmt"

const (
	p4_00 = iota << 4
	p4_01
	_
	_
	p4_02

	_
	_
	p4_03
)

func main() {
	fmt.Println(p4_00, p4_01, p4_02, p4_03)
}
