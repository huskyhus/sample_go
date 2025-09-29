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

	// simple if statement
	n := 0
	if n > 0 {
		fmt.Println("positive")
	} else if n < 0 {
		fmt.Println("negative")
	} else {
		fmt.Println("n == 0")
	}

	// define local variable
	n = 15
	if k := n * n; k > 200 {
		fmt.Println(k, "is greater than 200")
	} else {
		fmt.Println(k, "is less than or equal to 200")
	}

	fmt.Println("If: end")
}

func ForStatement() {
	fmt.Println("For: begin")

	// standard for statement
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// condition only (=while)
	i := 1
	for i < 100 {
		i *= 2
		fmt.Println(i)
	}

	// infinite loop with break
	i = 1
	for {
		if i%15 == 0 {
			fmt.Println(i, "FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println(i, "Fizz")
		} else if i%5 == 0 {
			fmt.Println(i, "Buzz")
		} else {
			fmt.Println(i)
		}

		i++
		if i > 100 {
			break
		}
	}

	// for-range for a slice
	slice := []int{100, 200, 300, 400, 500}
	for i, v := range slice {
		fmt.Println("slice[", i, "] is", v)
	}

	// for-range for a map
	dict := map[string]string{
		"apple":      "ringo",
		"banana":     "banana",
		"watermelon": "suika",
	}
	for k, v := range dict {
		fmt.Println("dict[", k, "] is", v)
	}

	// for-range for a string
	str := "こんにちは, 世界"
	for i, r := range str {
		fmt.Println(i, r, string(r))
	}
	// len(str) returns the number of bytes, not runes.
	// for-range loop for a string runs through runes.
	// thus, the following loop don't work.
	fmt.Println(len(str))
	for i := 0; i < len(str); i++ {
		fmt.Println(i, str[i])
	}

	fmt.Println("For: end")
}
