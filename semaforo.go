package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type register struct {
	ID   uint64
	Name string
}

var (
	wp sync.WaitGroup
	mx sync.Mutex
)

func main() {

	semaforo := make(chan struct{}, 10)

	for {
		wp.Add(2)

		semaforo <- struct{}{}
		semaforo <- struct{}{}

		go func(semaforo <-chan struct{}) {
			defer wp.Done()
			fmt.Println(fmt.Sprintf("semaforo count: %d", len(semaforo)))
			mx.Lock()
			defer mx.Unlock()
			err := conditionalError()
			if err != nil {
				fmt.Println("deu erro aqui.")
				<-semaforo
				return
			}
			<-semaforo
		}(semaforo)

		go func(semaforo <-chan struct{}) {
			defer wp.Done()
			fmt.Println(fmt.Sprintf("semaforo count: %d", len(semaforo)))
			err := conditionalError()
			if err != nil {
				fmt.Println("deu erro aqui.")
				<-semaforo
				return
			}
			<-semaforo
		}(semaforo)

		wp.Wait()

	}
}

func conditionalError() error {
	time.Sleep(500 * time.Millisecond)
	numberRandom := rand.Int()
	if numberRandom%2 == 0 {
		return errors.New("this number is pair")
	}
	return nil
}
