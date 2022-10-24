package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "20:22:33, 20:22:34"
	sl := strings.Split(s, ",")
	fmt.Println(strings.Join(sl[1:], "||"))
}
