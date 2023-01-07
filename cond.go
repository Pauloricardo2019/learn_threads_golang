package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	tam = 10
)

var (
	values [tam]string
)

func main() {

	rand.Seed(time.Now().UnixNano())
	
	cond := sync.NewCond(&sync.Mutex{})

	for i := 0; i < tam; i++ {
		d := time.Second * time.Duration(rand.Intn(10)) / 10
		go func(i int) {
			time.Sleep(d)
			cond.L.Lock()
			values[i] = string('a' + i)
			cond.L.Unlock()
			cond.Signal()
		}(i)
	}

	cond.L.Lock()
	defer cond.L.Unlock()
	for !checkCondition() {
		cond.Wait()
	}
 fmt.Println("finish process!")
}

func checkCondition() bool {
	fmt.Println(values)
	for i := 0; i < tam; i++ {
		if values[i] == "" {
			return false
		}
	}
	return true
}
