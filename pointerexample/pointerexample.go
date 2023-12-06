package pointerexample

import "fmt"

func Testintegerpointer() {
	x := 5
	fmt.Println("Mofifying x directly: ", modifyX(x))
	fmt.Println("after modification x directly: ", x)

	fmt.Println("X is: ", x)
	fmt.Println("X address is: ", &x)
	//pointer
	y := &x
	fmt.Println("Y is: ", y)

	fmt.Printf("Before modification %d\n", x)
	modify(y)
	fmt.Printf("After modification %d\n", x)

}

func modifyX(i int) int {
	i = 9
	return i
}

func modify(y *int) *int {
	*y = 10
	return y
}
