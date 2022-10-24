package main

import "fmt"

func printSlice(x []int) {
	for _, number := range x {
		fmt.Print(number, " ")
	}
	fmt.Println()
}
func printSliceInfo(x []int) {
	fmt.Printf("Cap: %d, Length: %d\n", cap(x), len(x))

}

func main() {
	aSlice := []int{-1, 0, 4}
	fmt.Print("aSlice: ")
	printSlice(aSlice)

	printSliceInfo(aSlice)
	aSlice = append(aSlice, -100)
	fmt.Printf("aSlice: ")
	printSlice(aSlice)
	printSliceInfo(aSlice)
	aSlice = append(aSlice, -2)
	aSlice = append(aSlice, -3)
	aSlice = append(aSlice, -4)
	printSlice(aSlice)
	printSliceInfo(aSlice)

}
