package main

import (
	"fmt"
	"leetcode/other/registration/process"
	"math"
	"time"
)

func main() {
	start := time.Now()
	registrants := 3

	for registrant := 1; registrant <= registrants; registrant++ {
		process.WriteName(registrant, 1)
		process.GetNameTag(registrant, 1)
		process.GetSnack(registrant, 1)
	}

	duration := time.Since(start)

	fmt.Println("Sequential: done in", int(math.Ceil(float64(duration.Milliseconds()))), "seconds")
}
