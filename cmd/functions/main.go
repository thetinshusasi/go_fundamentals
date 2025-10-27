package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func returnMulVals() (int, int) {
	return 3, 5
}

func returnMulVals2(a int, b int) (x int, y int) {
	x = a + 2
	y = b + 4
	return x, y
}

func main() {
	fmt.Println(" functions desc================")

	fmt.Println(add(100, 37))

	_, secVal := returnMulVals()
	fmt.Printf("SecondVal %v \n", secVal)

	_, secVal = returnMulVals2(10, 34)
	fmt.Printf("SecondVal %v \n", secVal)
}
