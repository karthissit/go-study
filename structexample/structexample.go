package structexample

import (
	"fmt"
)

type Person interface {
	Describe() string
}

type Employee struct {
	Id               int
	Name             string
	Designation      string
	PresentationName string
}

type Student struct {
	Id             int
	Name           string
	Class          string
	AggregateScore float32
}

func (e Employee) Describe() string {
	message := fmt.Sprintf("Employee name: %v, Employee Designation: %v, Employee Id: %v", e.Name, e.Designation, e.Id)
	return message
}

func (s Student) Describe() string {
	message := fmt.Sprintf("Student name: %v, Student Class: %v, Student Id: %v", s.Name, s.Class, s.Id)
	return message
}

func (e Employee) PresentSomething() {
	fmt.Printf("Employee %v is presenting %v\n", e.Name, e.PresentationName)
}

func (s Student) DisplayScore() {
	fmt.Printf("Student %v's aggregate score is %0.2f\n", s.Name, s.AggregateScore)
}
