package ds

import (
	"fmt"
	"strings"
)

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

	//slice

	numberSlice := []int16{}

	var i int16
	for i = 0; i < 10; i++ {
		numberSlice = append(numberSlice, i)
	}

	fmt.Println(numberSlice)

	//String slice

	stringSlice := []string{}

	for i := 0; i < 10; i++ {
		stringSlice = append(stringSlice, fmt.Sprint("String", i))
	}

	fmt.Println(stringSlice)

	//maps

	contryCityMap := make(map[string]string)

	contryCityMap["India"] = "New Delhi"
	contryCityMap["Sri Lanka"] = "Colombo"
	contryCityMap["Nepal"] = "Katmandu"
	contryCityMap["USA"] = "Washington DC"
	contryCityMap["Mexico"] = "Mexico City"

	fmt.Println(contryCityMap)

	for country, city := range contryCityMap {
		fmt.Printf("Capital City of the Country %v is %v\n", country, city)
	}

	//Challenge

	/*
		findDuplicatedString from following string using maps

		givenText := `Go is an open-source programming language created by Google,
		prized for its concurrency and connectivity.
		Using Go, developers can build modern applications that can save companies money on backend resources.
		This course is designed to help developers be productive with Go, starting with the essentials of the syntax.
		Learn the basics of Go basic types such as numbers and strings; working with conditionals and loops;
		creating object-oriented code with structs and methods; and handling errors.
		Instructor Miki Tebeka also emphasizes the concurrency features such as goroutines and channels,
		and connectivity features for networking with APIs and databases.
		For the final project, Mika shows you how to build a highly concurrent server that combines
		 everything you've learned into one elegant solution powered by Go.`
	*/

	givenText := `Go is an open-source programming language created by Google,
		prized for its concurrency and connectivity.
		Using Go, developers can build modern applications that can save companies money on backend resources.
		This course is designed to help developers be productive with Go, starting with the essentials of the syntax.
		Learn the basics of Go basic types such as numbers and strings; working with conditionals and loops;
		creating object-oriented code with structs and methods; and handling errors.
		Instructor Miki Tebeka also emphasizes the concurrency features such as goroutines and channels,
		and connectivity features for networking with APIs and databases.
		For the final project, Mika shows you how to build a highly concurrent server that combines
		 everything you've learned into one elegant solution powered by Go.`

	words := strings.Fields(givenText)
	wordsmap := map[string]int{}

	for _, word := range words {
		if wordsmap[word] == 0 {
			wordsmap[word] = 1
		} else {
			wordsmap[word] += 1
		}
	}

	fmt.Println(wordsmap)

	for word, repeated := range wordsmap {
		if repeated > 1 {
			fmt.Printf("The word %v is repated %d times\n", word, repeated)
		}

	}
}
