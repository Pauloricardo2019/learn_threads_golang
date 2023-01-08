package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	ch := make(chan int)

	go func() {

		sum := rand.Int() + rand.Int()

		ch <- sum
	}()

	result := <-ch

	fmt.Println("Result: ", result)
  fmt.Println("finish channel")
}
