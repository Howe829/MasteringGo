package main

import "fmt"

func main() {
	aByteSlice := make([]byte, 5)
	fmt.Println(aByteSlice)
	aByteSlice = append(aByteSlice, 10)
	fmt.Println(aByteSlice)

}
