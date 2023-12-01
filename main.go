package main

import (
	"fmt"

	"example.com/greeting"
)

func main() {
	message, err := greeting.Hello("John")
	if err != nil {
		fmt.Println("Error in input", err)
	}
	fmt.Println(message)
}
