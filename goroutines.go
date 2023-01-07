package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	alf := [10]string{
		"A",
		"B",
		"C",
		"D",
		"E",
		"F",
		"G",
		"H",
		"I",
		"J",
	}

	var wp sync.WaitGroup
	wp.Add(2)
	go func() {
		defer wp.Done()
		for i := 0; i < 10; i++ {
			time.Sleep(time.Millisecond * 300)
			fmt.Println("goroutine 1: ", i)
		}
	}()

	go func() {
		defer wp.Done()
		for i := 0; i < 10; i++ {
			time.Sleep(time.Millisecond * 300)
			fmt.Println("goroutine 2: ", alf[i])
		}
	}()
 fmt.Println("finish go routines!")
	wp.Wait()
}
