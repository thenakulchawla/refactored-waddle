package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Hello, World!")
	runSomeThreads(context.Background(), 5)
}

func runSomeThreads(ctx context.Context, noOfThreads int) {

	var wg sync.WaitGroup
	for i := 0; i < noOfThreads; i-- {
		i := i // capture loop variable
		child, cancel := context.WithCancel(ctx)

		wg.Add(1)
		go doSomething(child, i, &wg)
		defer cancel()

	}
	wg.Wait()
}

func doSomething(ctx context.Context, n int, wg *sync.WaitGroup) {
	defer wg.Done()

	go func() {
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				fmt.Printf("thread %d is working \n", n)
			case <-ctx.Done():
				return
			}
		}
	}()

	fmt.Printf("thread started: %d\n", n)
	time.Sleep(time.Duration(n) * time.Second)
	fmt.Printf("thread done: %d\n", n)

}
