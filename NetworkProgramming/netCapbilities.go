package main

import (
	"fmt"
	"net"
)

func main() {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Print(err)
	}
	for _, i := range interfaces {
		fmt.Printf("Name:%v\n", i.Name)
		fmt.Printf("Interface Flags:%v\n", i.Flags)

		fmt.Printf("Internet MTU:%v\n", i.MTU)

		fmt.Printf("Internet Hardware Address:%v\n", i.HardwareAddr)
		fmt.Println()

	}
}
