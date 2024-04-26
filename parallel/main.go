package main

import (
	"bytes"
	"fmt"
	"leetcode/other/registration/process"
	"math"
	"runtime/debug"
	"strconv"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	registrants := 3

	wg := new(sync.WaitGroup)

	queue := make(chan int)

	go doTheJobs(queue, wg)

	for registrant := 1; registrant <= registrants; registrant++ {
		wg.Add(1)
		queue <- registrant
	}

	wg.Wait()

	duration := time.Since(start)

	fmt.Println("Parallel: done in", int(math.Ceil(float64(duration.Milliseconds()))), "seconds")
}

func doTheJobs(queue chan int, wg *sync.WaitGroup) {
	for i := 1; i <= 2; i++ {
		go func(i int) {
			gr := bytes.Fields(debug.Stack())[1]
			grString := string(gr)
			index, _ := strconv.Atoi(grString)
			for registrant := range queue {
				process.WriteName(registrant, index)
				process.GetNameTag(registrant, index)
				process.GetSnack(registrant, index)
				wg.Done()
			}
		}(i)
	}

}
