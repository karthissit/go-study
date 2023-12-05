package pointerexample

import "fmt"

func Testintegerpointer() {

	x := 5

	y := &x

	fmt.Println("Mofifying x directly: ", modifyX(x))
	fmt.Println("after modification x directly: ", x)

	fmt.Println("X is: ",x)
	fmt.Println("X address is: ",&x)
	fmt.Println("Y is: ",y)

	fmt.Printf("Before modification %d\n", x)
	modify(y)
	fmt.Printf("After modification %d\n", x)

}

func modifyX(x int) int {
	x = 9
	return x
}

func modify(y *int) *int {
	*y = 10
	return y
}
