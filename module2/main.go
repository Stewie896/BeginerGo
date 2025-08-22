package main

import (
	"fmt"
)

func main() {

	counterArr := [...]int{34, 35, 3, 4, 34}

	for i := 0; i <= len(counterArr)-1; i++ {
		fmt.Println(counterArr[i])
		if counterArr[i] == 3 {
			fmt.Println("The number ", counterArr[i], " is present")
		}
	}

	counter := 10

	switch counter {
	case 1:
		fmt.Println("")
	}

	newArr := make([]string, 0)
	newArr = append(newArr, "Yalon")
	fmt.Println(len(newArr))
	fmt.Println(cap(newArr), " this is the ttal capacity that can fit/.")

	//nil -> ....if(err != nil)

	keyVal := make(map[string]float32)
	keyVal["name"] = 343.33
	keyVal["age"] = 17.9
	valueReveived, ok := keyVal["name"]

	if ok {
		fmt.Println("Value reveived ", valueReveived)
	} else {
		fmt.Println("The value that you specidfied is not present")
	}

	// constructArr := make([]string, 0, 10)
	// //onstructArr[0] = "hello"
	// constructArr = append(constructArr, "Abhiral")

	// sum := 0

	// for i := 0; i <= len(counterArr)-1; i++ {
	// 	sum += 1
	// }
	// fmt.Println(sum, " is the value ")

	for key, val := range keyVal {
		fmt.Println("Hello iam key ", key)
		fmt.Println("Hello iam value ", val)
	}

}

//console.log("Hello world");
// Logical scripting
