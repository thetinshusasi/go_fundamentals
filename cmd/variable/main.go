package main

import "fmt"

func main() {

	var intVal int
	var floatVal float64
	var boolVal bool
	var stringVal string

	fmt.Printf("%v %f %v %q \n", intVal, floatVal, boolVal, stringVal)

	intVal2 := 100

	congrats := " happy birthday"

	fmt.Printf("%v %s \n", intVal2, congrats)

	fmt.Printf("%T %T %T %T \n", intVal, floatVal, boolVal, stringVal)

	fmt.Println(intVal, "sdafsdf ", intVal2)

	// constants

	const constVal = "This is a constant"
	// constVal = " this won't work"
	fmt.Printf("%s \n", constVal)

	if true {
		fmt.Println("conditonal examples")
	}
}
