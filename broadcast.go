package main

import (
	"fmt"
	"sync"
	"time"
)

var sharedRsc = false

func main() {
	var wg sync.WaitGroup

	wg.Add(4)
	c := sync.NewCond(&sync.Mutex{})

	go func() {
		c.L.Lock()
		for sharedRsc == false {
			fmt.Println("goroutine1 wait")
			c.Wait()
		}
		fmt.Println("goroutine1 ", sharedRsc)
		fmt.Println(time.Now().Format("15:04:05,000"))
		fmt.Println("--------------------- ")

		c.L.Unlock()
		wg.Done()
	}()

	go func() {
		c.L.Lock()
		for sharedRsc == false {
			fmt.Println("goroutine2 wait")
			c.Wait()
		}

		fmt.Println("goroutine2 ", sharedRsc)
		fmt.Println(time.Now().Format("15:04:05,000"))
		fmt.Println("--------------------- ")
		c.L.Unlock()
		wg.Done()
	}()

	go func() {
		c.L.Lock()
		for sharedRsc == false {
			fmt.Println("goroutine3 wait")
			c.Wait()
		}
		fmt.Println("goroutine3 ", sharedRsc)
		fmt.Println(time.Now().Format("15:04:05,000"))
		fmt.Println("--------------------- ")

		c.L.Unlock()
		wg.Done()
	}()

	go func() {
		c.L.Lock()
		for sharedRsc == false {
			fmt.Println("goroutine4 wait")
			c.Wait()
		}
		fmt.Println("goroutine4 ", sharedRsc)
		fmt.Println(time.Now().Format("15:04:05,000"))
		fmt.Println("--------------------- ")

		c.L.Unlock()
		wg.Done()
	}()

	time.Sleep(time.Second * 2)

	c.L.Lock()
	fmt.Println("main goroutine ready")
	sharedRsc = true
	c.Broadcast()
	fmt.Println("main goroutine broadcast")
	c.L.Unlock()

	wg.Wait()

}
