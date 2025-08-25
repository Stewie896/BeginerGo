package main

import (
	"fmt"
	"runtime"
	"sync"
)

func iterate(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= 9_000_000_000; i++ {
		//time.Sleep(10 * time.Millisecond)
	}
	fmt.Println("1 billon done on iteration function")
}

func detrate(wg *sync.WaitGroup) {
	defer wg.Done()
	for x := 11; x <= 9_000_000_000; x++ {
		//time.Sleep(10 * time.Millisecond)
	}
	fmt.Println("1 billlon done on  deiterate func")
}

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	runtime.GOMAXPROCS(16)
	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0))

	var wg sync.WaitGroup
	wg.Add(2)

	go iterate(&wg)
	go detrate(&wg)

	fmt.Println("Concurrency in action")

	wg.Wait()
}
