package main

import (
		"net"
		"fmt"
)

func main() {
	lis, err := net.Listen("tcp","localhost:4000")
	if err!=nil {
		fmt.Println("err =", err)
	}
}