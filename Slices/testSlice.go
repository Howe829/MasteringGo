package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := [5]string{"1", "2", "3", "4", "4"}
	dup := make([]string, 0)
	for _, n := range a {
		if !InSlice(n, dup) {
			dup = append(dup, n)
		}
	}
	fmt.Println(dup)
}

func InSlice(v interface{}, sl interface{}) bool {
	ret := reflect.ValueOf(sl)
	if ret.Kind() != reflect.Slice {
		panic("InterfaceSlice() given a non-slice type")
	}

	// Keep the distinction between nil and empty slice input
	if ret.IsNil() {
		return false
	}
	for i := 0; i < ret.Len(); i++ {
		if v == ret.Index(i).Interface() {
			return true
		}
	}
	return false
}
