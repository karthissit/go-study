package main

import (
	"fmt"

	//user defined
	"example.com/conditionals"
	"example.com/ds"
	"example.com/greeting"
	"example.com/loops"
	"example.com/numbers"
	"example.com/str"
)

func main() {

	//Anatomy
	message, err := greeting.Hello("John")
	if err != nil {
		fmt.Println("Error in input", err)
	}
	fmt.Println(message)

	//numbers
	numbers.Numbers()

	//string
	str.StringExample()

	//conditionals
	conditionals.ConditionalsExample(10, 12)

	//loops
	loops.LoopsExample()

	//ds
	ds.Datastructures()

}
