package main

import "fmt"

type rect struct {
	width  float64
	height float64
}

func main() {

	// hits := make(map[string]map[string]int)

	ages := map[string]int{
		"tinshu": 10,
		"titu":   20,
	}

	fmt.Println(ages)

	if elm, ok := ages["titu"]; !ok {
		fmt.Println("elm not found")
	} else {
		fmt.Println(elm)
	}

	delete(ages, "titu")

	if elm, ok := ages["titu"]; !ok {
		fmt.Println("elm not found")
	} else {
		fmt.Println(elm)
	}

	//we cann't use slices, maps and function as keys in a map
	// but can use anything else for example we can use struct as  key

	rect1 := rect{
		width:  10,
		height: 30,
	}

	//
	structAsKey := map[rect]int{
		rect1: 30,
	}

	fmt.Println(structAsKey)

}
