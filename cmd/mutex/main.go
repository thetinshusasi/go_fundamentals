package main

import (
	"fmt"
	"sync"
	"time"
)

type connect interface {
	myconnection()
}

type test interface {
	mytest(t connect)
}

type test2[c connect] interface {
	mytest(c)
}

var wg *sync.WaitGroup

func dowork(delay time.Duration, resultch chan string) {
	defer wg.Done()
	time.Sleep(delay)
	resultch <- fmt.Sprintf("Done in %s", delay)
	fmt.Println("dowork Done in ", delay)

}

func main() {
	start := time.Now()
	/// we need a buffered channel , if use unbuffered channel,  3 routine are waiting to write to resultch,
	resultch := make(chan string)
	wg = &sync.WaitGroup{}
	wg.Add(3)
	go dowork(time.Second*2, resultch)
	go dowork(time.Second*2, resultch)
	go dowork(time.Second*2, resultch)

	// or this way
	// go func() {
	// 	wg.Wait()
	// 	close(resultch)
	// }()

	// for res := range resultch {
	// 	fmt.Println(res)
	// }

	go func() {
		for res := range resultch {
			fmt.Println(res)
		}

	}()

	wg.Wait()

	close(resultch)

	fmt.Printf("Total time: %s\n", time.Since(start))

}
