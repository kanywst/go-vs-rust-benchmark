package main

import (
	"fmt"
	"runtime"
	time "time"
)

type Payload struct {
	data [1024]byte
}

var Sink *Payload

func main() {
	count := 10_000_000 
	start := time.Now()
	
	for i := 0; i < count; i++ {
		p := &Payload{}
		Sink = p 
	}

	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	
	fmt.Printf("Go: Allocations took %v\n", time.Since(start))
	fmt.Printf("Go: NumGC = %v\n", m.NumGC)
}
