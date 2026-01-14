package main

import (
	"fmt"
	time "time"
)

func main() {
	size := 100000000
	data := make([]int, size)
	for i := 0; i < size; i++ { data[i] = i }

	start := time.Now()
	sum := 0
	for _, v := range data {
		sum += v
	}

	fmt.Printf("Go: Sum %d in %v\n", sum, time.Since(start))
}
