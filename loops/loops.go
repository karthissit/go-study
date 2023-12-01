package loops

import "fmt"

func LoopsExample() {

	//only one loop in go lang "for" ==> can be used as while, do while, for, for each

	// regular
	count := 10
	for i := 0; i < count; i++ {
		fmt.Printf("In loop %d\n", i)
	}

	//while
	for 0 < count {
		fmt.Printf("Decremented to: %d", count)
		count--
	}

	//do while
	b := 0
	for {
		if b > 2 {
			break
		}
		fmt.Println(b)
		b++
	}

	//for each
	intslice := []int{1, 2, 3, 4, 5}
	for i, num := range intslice {
		fmt.Printf("Index: %d, Number: %d", i, num)
	}
}
