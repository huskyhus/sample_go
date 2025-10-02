package learning_go

import (
	"fmt"
	"io"
	"os"
)

// variadic parameters
func sum(nums ...int) (total int) {
	for _, n := range nums {
		total += n
	}
	return total
}

// function type
type binaryOperator func(float64, float64) float64

func FunctionExample() {
	fmt.Println("FunctionExample: begin")

	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(sum([]int{6, 7, 8, 9, 10}...))

	// a function is an object
	var operators = map[string]binaryOperator{
		"add":      func(a, b float64) float64 { return a + b },
		"subtract": func(a, b float64) float64 { return a - b },
		"multiply": func(a, b float64) float64 { return a * b },
		"divide":   func(a, b float64) float64 { return a / b },
	}
	for _, operator := range operators {
		fmt.Println(operator(10, 5))
	}

	fmt.Println("FunctionExample: end")
}

func Cat(filename string) {
	// open the file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// read the file by chunks
	data := make([]byte, 64)
	for {
		count, err := file.Read(data)
		if err != nil && err != io.EOF {
			fmt.Println("Error:", err)
			return
		} else if err == io.EOF {
			break
		} else {
			fmt.Print(string(data[:count]))
		}
	}
}

// closer pattern
func getFile(filename string) (*os.File, func(), error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}

	return file, func() {
		file.Close()
	}, nil
}

func Cat2(filename string) {
	file, closer, err := getFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer closer()

	// read the file by chunks
	data := make([]byte, 64)
	for {
		count, err := file.Read(data)
		if err != nil && err != io.EOF {
			fmt.Println("Error:", err)
			return
		} else if err == io.EOF {
			break
		} else {
			fmt.Print(string(data[:count]))
		}
	}
}
