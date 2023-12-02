package str

import "fmt"

func StringExample() {
	myname := "Karthikeyan"

	fmt.Println(myname)

	multilineString := `Designed to run on a variety of footprints 
	that range from a traditional on-premises data center 
	to a public cloud environment.`

	fmt.Println(multilineString)

	//Processing on Strings

	fmt.Printf("Printing within range 0 to 6 %v\n", myname[0:6])
	fmt.Printf("Printing within range until 6 %v\n", myname[:6]) //result is same as above

	fmt.Printf("Printing within range 6 to end %v\n", myname[6:])
	fmt.Printf("Printing within range 3 to 7 %v\n", myname[3:7])
}
