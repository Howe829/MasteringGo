package main

import "fmt"

func main() {
	myArray := [5]int{1, 2, 3, 4, 5}
	myMap := make(map[int]int)
	for i, j := range myArray {
		myMap[i] = j
	}
	fmt.Println(myMap)
}
