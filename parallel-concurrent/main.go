package main

import (
	"fmt"
	"leetcode/other/registration/process"
	"math"
	"sync"
	"time"
)

func main() {
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

	fmt.Println("Parallel concurrent: done in", int(math.Ceil(float64(duration.Milliseconds()))), "seconds")
}

func createRegistrationTables(queue chan int, wg *sync.WaitGroup) {
	nameTagQueue := make(chan int)
	snackQueue := make(chan int)

	for i := 1; i <= 2; i++ {
		go func(i int) {
			for registrant := range queue {
				process.WriteName(registrant, i)
				nameTagQueue <- registrant
			}
		}(i)

		go func(i int) {
			for registrant := range nameTagQueue {
				process.GetNameTag(registrant, i)
				snackQueue <- registrant
			}
		}(i)

		go func(i int) {
			for registrant := range nameTagQueue {
				process.GetNameTag(registrant, i)
				snackQueue <- registrant
			}
		}(i)

		go func(i int) {
			for registrant := range snackQueue {
				process.GetSnack(registrant, i)
				wg.Done()
			}
		}(i)

	}

}
