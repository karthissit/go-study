package ds

import "fmt"

func Datastructures() {
	// arrays -> fixed size
	// slice -> equivalent to list or not a fixed array
	// map -> map of key value pairs

	//array
	var numberArray [5]int

	numberArray[0] = 10

	// default the primitive types or the constraints.Ordered types will be initalized with the default 0
	/* OUTPUT
	10
	0
	0
	0
	0
	*/

	for i := 0; i < len(numberArray); i++ {
		fmt.Println(numberArray[i])
	}
}
