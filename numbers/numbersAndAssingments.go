package numbers

import "fmt"

func Numbers() {
	/*
		int, int8, int16, int32, int64,

		uint, uint8, uint16, uint32, uint64

		float32, float64

	*/
	var x int
	var y int

	x = 5
	y = 6

	fmt.Printf("x: %d, y: %d\n", x, y)

	a := 10
	b := 12

	fmt.Printf("a: %d, b: %d\n", a, b)

	var f1 float32
	var f2 float64

	f1 = 10.2
	f2 = 11.2

	fmt.Printf("f1: %f, f2: %f\n", f1, f2)
}
