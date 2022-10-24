package main

import "fmt"

func main() {
	iMap := make(map[string]int)
	iMap["k1"] = 12
	iMap["k2"] = 13
	fmt.Println("iMap:", iMap, iMap["k3"])
	anotherMap := map[string]int{
		"k1": 14,
		"k2": 15,
	}
	fmt.Println(anotherMap)
	_, ok := iMap["DoesNotExist"]
	if ok {
		fmt.Println("Exists!")
	} else {
		fmt.Println("Does not exist!")
	}
	for key, value := range iMap {
		fmt.Println("Key: ", key, "Value: ", value)
	}
	delete(anotherMap, "k1")
	for key, value := range anotherMap {
		fmt.Println("Key: ", key, "Value: ", value)
	}
}
