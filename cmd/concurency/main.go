package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func readFromChannel(ch <-chan int) {
	// Loop over the channel until it's closed
	for num := range ch {
		fmt.Println("Received readFromChannel:", num)
	}
}

func writeToChannel(ch chan<- int) {
	// Send some integers to the channel
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch) // Close the channel after sending all values
}

func main() {

	// for i := 0; i < 10; i++ {

	// 	go func() {
	// 		time.Sleep(time.Millisecond * 250)
	// 		fmt.Println("Email Recieved")
	// 	}()

	// 	fmt.Println("Email Send")
	// }

	chInt := make(chan int, 4)

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("data sent", i)

			chInt <- i
		}
		// close(chInt)
	}()

	// for val := range chInt {
	// 	fmt.Println("i : ", val)

	// }

	fmt.Println("i : ", <-chInt)
	fmt.Println("i : ", <-chInt)
	time.Sleep(time.Millisecond * 250)
	fmt.Print("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	fmt.Println("i : ", <-chInt)

	ch1 := make(chan string)
	ch2 := make(chan string)

	// Goroutine sending "message 1" to ch1 after 2 seconds
	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "message 1"
	}()

	// Goroutine sending "message 2" to ch2 after 3 seconds
	go func() {
		time.Sleep(3 * time.Second)
		ch2 <- "message 2"
	}()

	// Select statement to wait for messages from multiple channels
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Received from ch1:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Received from ch2:", msg2)
		}
	}

	ch := make(chan int)

	go func() {
		writeToChannel(ch)
	}()

	// Function to read from the read-only channel
	readFromChannel(ch)

}
