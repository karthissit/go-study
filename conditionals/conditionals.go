package conditionals

import "fmt"

func ConditionalsExample(x int, y int) {

	//ifelse
	if x > y {
		fmt.Print("x is gt than y")
	} else {
		fmt.Print("y is gt than x")
	}

	if fraction := x / y; fraction > 1 {
		fmt.Printf("The fraction %v is greater than 1\n", fraction)
	}

	//switch

	n := 10

	switch n {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("TWO")
	/*
		Skipped 3 - 9
	*/
	case 10:
		fmt.Println("TEN")
	default:
		fmt.Println("OUT OF RANGE")
	}

	switch {
	case n == 5:
		fmt.Println("Equals 5")
	case n > 5:
		fmt.Println("Greater than 5")
	case n < 5:
		fmt.Println("Less than 5")
	case n == 10:
		fmt.Println("Is 10")
	default:
		fmt.Println("OUT OF RANGE")
	}

}
