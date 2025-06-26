package main

import (
	"runtime"
	"sync"
)

func main() {
	// Additional parallelizer.
	runtime.GOMAXPROCS(1000)

	wg := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			singleton := Instance()
			singleton.DoJob()
		}()
	}

	wg.Wait()
}
