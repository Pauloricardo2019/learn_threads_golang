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
)

func main() {

	semaforo := make(chan struct{}, 10)

	for {
		wp.Add(1)

		semaforo <- struct{}{}

		go func(semaforo <-chan struct{}) {

			defer wp.Done()

			fmt.Println(fmt.Sprintf("semaforo position: %d", len(semaforo)))

			err := conditionalError()
			if err != nil {
				fmt.Println("deu erro aqui.")
				<-semaforo
				return
			}
			<-semaforo

		}(semaforo)

		wp.Wait()

		time.Sleep(time.Millisecond * 300)
	}
}

func conditionalError() error {
	numberRandom := rand.Int()
	if numberRandom%2 == 0 {
		return errors.New("this number is pair")
	}
	return nil
}
