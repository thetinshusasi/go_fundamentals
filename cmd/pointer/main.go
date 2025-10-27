package main

import "fmt"

// When to Use Pointer Receivers
// Use pointer receivers when:

// The method needs to modify the receiver.
// The receiver is a large struct and copying would be inefficient.
// You want to ensure that all methods on the type have consistent receivers.
// Value vs. Pointer Receivers
// Value Receiver: A method with a value receiver cannot modify the original struct and works on a copy of the receiver.
// Pointer Receiver: A method with a pointer receiver can modify the original struct.

type Counter struct {
	count int
}

func (c *Counter) IncrementPtr() {
	c.count++
	fmt.Println("ptr")
}
func (c Counter) IncrementVal() {
	c.count++
	fmt.Println("Val")
}

func main() {
	c := Counter{}
	fmt.Println("Counter : ", c.count)

	c.IncrementPtr() // Go converts this to (&c).IncrementPtr()
	fmt.Println("Counter : ", c.count)

	(&c).IncrementPtr() // Explicitly calling with a pointer
	fmt.Println("Counter : ", c.count)

	c.IncrementVal() // Calls normally
	fmt.Println("Counter : ", c.count)

	(&c).IncrementVal() // is invalid because it requires a value receiver

	fmt.Println("Counter : ", c.count)

}
