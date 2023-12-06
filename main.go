package main

import (
	"fmt"

	//user defined
	"example.com/conditionals"
	"example.com/ds"
	"example.com/greeting"
	"example.com/loops"
	"example.com/numbers"
	"example.com/pointerexample"
	"example.com/str"
	"example.com/structexample"
	"rsc.io/quote"
)

const PI = 3.14

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

	fmt.Printf("DEFINED CONSTANT: PI is %f\n", PI)

	//Using external modules
	fmt.Println(quote.Glass())
	fmt.Println(quote.Hello())
	fmt.Println(quote.Opt())
	fmt.Println(quote.Go())

	//test pointers
	pointerexample.Testintegerpointer()

	//Structures and interfaces
	emp1 := structexample.Employee{
		Name:             "Karthikeyan",
		Designation:      "TS",
		Id:               1234,
		PresentationName: "GoLang",
	}

	stu1 := structexample.Student{
		Name:           "Khatvik",
		Class:          "LKG",
		Id:             123123,
		AggregateScore: 87.23,
	}

	empMessage := emp1.Describe()
	stuMessage := stu1.Describe()

	fmt.Printf("Employee Message: %v\n", empMessage)
	fmt.Printf("Student Message: %v\n", stuMessage)

	emp1.PresentSomething()
	stu1.DisplayScore()

}
