package main

import (
	"fmt"
	"runtime"
	"sync"
)

func task1() {
	for i := 0; i < 10; i++ {
		fmt.Println(i, "Hellow")
	}
	wg.Done()
}

func task2() {
	for i := 0; i < 10; i++ {
		fmt.Println(i, "World")
	}
	wg.Done()
}

var wg sync.WaitGroup

func main() {

	// time.Sleep(time.Second) // Wait to let goroutines complete

	
	wg.Add(1)
	go task1() // Runs concurrently
	go task2() // Runs concurrently
	wg.Wait()


	fmt.Println("Number of cpu", runtime.NumCPU())
	fmt.Println("Number of gorutiens", runtime.NumGoroutine())
}
