package learning_go

import (
	"fmt"
	"math"
)

func BasicTypes() {
	fmt.Println("Basic Types: begin")

	// bool
	var b = true
	fmt.Println(b && false)
	fmt.Println(b || false)
	fmt.Println(!b)

	// int
	var i = 12345
	fmt.Println(i + 100)
	fmt.Println(i - 100)
	fmt.Println(i * 100)
	fmt.Println(i / 100)

	// float64
	var f = 3.1415926535
	fmt.Println(math.Pow(f, 100))
	fmt.Println(math.Pow(f, 0.01))

	// rune = int32
	var r = 'A'
	fmt.Println(r)

	// string
	var s = "Hello, World!"
	fmt.Println(s)

	// byte = uint8
	var char = s[3]
	fmt.Println(char)

	// some variants to declare variables
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

	// constant
	const constant int = 42
	// constant += 100
	fmt.Println(constant)

	fmt.Println("Basic Types: end")
}
