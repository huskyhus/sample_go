package learning_go

import (
	"fmt"
)

func CompoundTypes() {
	fmt.Println("Compound Types: begin")

	// array
	var array = [3]int{1, 2, 3}
	var array2 = [...]float64{1.1, 2.2, 3.3, 4.4, 5.5}
	fmt.Println(array, len(array), cap(array))
	fmt.Println(array2, len(array2), cap(array2))

	// slice
	var slice = []int{4, 5, 6, 7}
	fmt.Println(slice, len(slice), cap(slice))
	slice = append(slice, slice...)
	slice = append(slice, 99, 100, 101)
	fmt.Println(slice, len(slice), cap(slice))
	fmt.Println(slice[3:5])

	// make a slice
	var maked = make([]int, 5, 12)
	fmt.Println(maked, len(maked), cap(maked))

	// copy a slice
	copy(maked, slice)
	fmt.Println(maked, len(maked), cap(maked))

	// map
	var m = map[string]int{
		"foo": 1,
		"bar": 2,
	}
	fmt.Println(m, len(m))
	m["baz"] = 3
	fmt.Println(m["foo"], m["bar"], m["baz"])

	// delete from a map
	delete(m, "foo")
	v, ok := m["foo"]
	fmt.Println(v, ok)

	// make a map
	var makedMap = make(map[string]int, 10)
	fmt.Println(makedMap, len(makedMap))

	// define a struct type
	type person struct {
		name string
		age  int
	}
	var p = person{
		name: "Alice",
		age:  30,
	}
	fmt.Println(p)
	fmt.Println(p.name, p.age)

	// make an anonymous struct
	var point = struct {
		x int
		y int
	}{640, 480}
	fmt.Println(point)

	fmt.Println("Compound Types: end")
}
