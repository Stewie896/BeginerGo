package main

import "fmt"

func receiver[X any](value []X) {
	ar := make([]X, len(value))
	for i := 0; i <= len(value)-1; i++ {
		ar[i] = value[i]
	}

}

func main() {
	arr := [2][3]int{
		{1, 2, 3},
		{2, 23, 4},
	}

	for _, value := range arr {
		for _, rlval := range value {
			fmt.Println(rlval)
		}
	}

	randomArr := [...]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	receiver(randomArr[:])

}
