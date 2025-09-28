package learning_go

import (
	"fmt"
)

func Shadowing() {
	fmt.Println("Shadowing: begin")

	x := 10
	fmt.Println(x)
	if x > 0 {
		x := -10
		fmt.Println(x)
	}
	fmt.Println(x)

	fmt.Println("Shadowing: end")
}

func IfStatement() {
	fmt.Println("If: begin")

	// TODO: fill this

	fmt.Println("If: end")
}