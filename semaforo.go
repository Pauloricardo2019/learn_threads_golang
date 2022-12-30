package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type register struct {
	ID   uint64
	Name string
}

var (
	wp sync.WaitGroup
)

func main() {

	semaforo := make(chan struct{}, 10)

	for {
		wp.Add(1)
		semaforo <- struct{}{}
		go func(semaforo <-chan struct{}) {
			defer wp.Done()
			fmt.Println(fmt.Sprintf("semaforo position: %d", len(semaforo)))
			_, err := generate()
			if err != nil {
				fmt.Println("deu erro aqui.")
				return
			}
			<-semaforo
		}(semaforo)
		wp.Wait()
		time.Sleep(time.Millisecond * 300)
	}
}

func generate() (*string, error) {
	return nil, errors.New("deu ruim")
}
