package main

import (
	"fmt"
	"strconv"
)

func add(vals []int) int {
	res := 0
	for i := 0; i < len(vals); i++ {
		res += vals[i]
	}
	return res
}
func add2(vals ...int) int {
	res := 0
	for i := 0; i < len(vals); i++ {
		res += vals[i]
	}
	return res
}

func main() {

	// int assigned to default value FOR ARRAYS
	var myInts [10]int

	for i := 0; i < len(myInts); i++ {
		fmt.Print(strconv.Itoa(myInts[i]) + " ")
	}
	fmt.Println("")
	primes := [4]int{2, 3, 5, 7}

	for i := 0; i < len(primes); i++ {
		fmt.Print(strconv.Itoa(primes[i]) + " ")
	}

	// slices
	fmt.Println("")

	primesSubs := primes[1:3]

	for i := 0; i < len(primesSubs); i++ {
		fmt.Print(strconv.Itoa(primesSubs[i]) + " ")

	}
	fmt.Println("")
	myInt2 := []int{}
	fmt.Printf("Capacity :  %v , Len : %v", cap(myInt2), len(myInt2))

	for i := 0; i < len(primesSubs); i++ {

		myInt2 = append(myInt2, primesSubs[i])

	}
	fmt.Println("")
	for i := 0; i < len(myInt2); i++ {

		fmt.Print(strconv.Itoa(myInt2[i]) + " ")

	}
	fmt.Printf("Capacity :  %v , Len : %v", cap(myInt2), len(myInt2))
	fmt.Println("")

	fmt.Println(add(primes[:2]))
	fmt.Println(add2(primes[:2]...))

	for ind, ele := range primes {
		fmt.Println(ind, ":", ele)

	}

}
