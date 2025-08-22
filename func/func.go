package main

import "fmt"

// Filling in the gaps
func Hello(value *int) {
	*value = 60
	fmt.Println("Is a modified ? ", *value)
}

func main() {
	value := 30
	Hello(&value)
	fmt.Println("This is the orginal value ", value)
}
