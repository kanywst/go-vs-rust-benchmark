package main

import (
	"fmt"
	"sync"
	time "time"
)

func main() {
	start := time.Now()
	var wg sync.WaitGroup
	count := 100000 // 10万goroutine

	wg.Add(count)
	for i := 0; i < count; i++ {
		go func() {
			defer wg.Done()
			time.Sleep(1 * time.Nanosecond) // 擬似的な処理
		}()
	}
	wg.Wait()
	fmt.Printf("Go: Spawning %d goroutines took %v\n", count, time.Since(start))
}
