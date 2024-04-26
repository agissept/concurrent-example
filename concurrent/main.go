package main

import (
	"fmt"
	"leetcode/other/registration/process"
	"math"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	start := time.Now()
	registrants := 3

	wg := new(sync.WaitGroup)

	queue := make(chan int)

	go createRegistrationTables(queue, wg)

	for registrant := 1; registrant <= registrants; registrant++ {
		wg.Add(1)
		queue <- registrant
	}

	wg.Wait()

	duration := time.Since(start)

	fmt.Println("Concurrent: done in", int(math.Ceil(float64(duration.Milliseconds()))), "seconds")
}

func createRegistrationTables(queue chan int, wg *sync.WaitGroup) {
	nameTagQueue := make(chan int)
	snackQueue := make(chan int)

	go func() {
		for registrant := range queue {
			process.WriteName(registrant, 1)
			nameTagQueue <- registrant
		}
	}()

	go func() {
		for registrant := range nameTagQueue {
			process.GetSnack(registrant, 1)
			snackQueue <- registrant
		}
	}()

	go func() {
		for registrant := range snackQueue {
			process.GetSnack(registrant, 1)
			wg.Done()
		}
	}()

}
