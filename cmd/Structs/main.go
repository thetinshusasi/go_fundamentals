package main

import (
	"fmt"
)

type Health struct {
	level int
}

type Player struct {
	Health
	name string
	age  int
}

type rect struct {
	width  float64
	height float64
}

func (r rect) area() float64 {
	return r.height * r.width
}

func main() {
	player := Player{
		name: "Tinshu",
		age:  12,
		Health: Health{
			level: 20,
		},
	}
	// empty struct
	player2 := Player{}

	fmt.Println(player.age, player.name, player.level)
	fmt.Println(player2.age, player2.name)

	rect1 := rect{
		width:  20,
		height: 4,
	}

	fmt.Println("area of rect ", rect1.area())

}
