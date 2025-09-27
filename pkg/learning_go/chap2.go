package learning_go

import (
	"fmt"
	"math"
)

func BasicTypes() {
	fmt.Println("Basic Types: begin")
	var b = true
	fmt.Println(b && false)
	fmt.Println(b || false)
	fmt.Println(!b)

	var i = 12345
	fmt.Println(i + 100)
	fmt.Println(i - 100)
	fmt.Println(i * 100)
	fmt.Println(i / 100)

	var f = 3.1415926535
	fmt.Println(math.Pow(f, 100))
	fmt.Println(math.Pow(f, 0.01))

	var r = 'A'
	fmt.Println(r)

	var s = "Hello, World!"
	fmt.Println(s)

	var (
		listed       = 10
		values       = "string"
		defaultValue int
	)
	fmt.Println(listed)
	fmt.Println(values)
	fmt.Println(defaultValue)

	shorthand := "shorthand"
	fmt.Println(shorthand)

	const constant int = 42
	// constant += 100
	fmt.Println(constant)

	fmt.Println("Basic Types: end")
}
