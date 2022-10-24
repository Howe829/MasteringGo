package main

import "fmt"

func getPointer(n *int) {
	*n = *n * *n
}

func returnPointer(n int) *int {
	v := n * n
	return &v
}

func main() {
	i := -10
	j := 25
	pI := &i
	pJ := &j

	fmt.Println("i Memory ", pI)
	fmt.Println("j Memory ", pJ)
	fmt.Println("pI Value ", *pI)

	fmt.Println("pJ Value ", *pJ)

	*pI = 123456
	*pI--
	fmt.Println("i: ", i)

	getPointer(pJ)
	fmt.Println("j: ", j)
	k := returnPointer(12)
	fmt.Println(*k)
	fmt.Println(k)
}
