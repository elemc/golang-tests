package main

import "fmt"

func testSlices1() {
	a := []string{"a", "b", "c"}
	b := a[1:2]
	b[0] = "q"
	fmt.Printf("%s\n", a)
}

func testSlices2() {
	a := []byte{'a', 'b', 'c'}
	b := append(a[1:2], 'd')
	b[0] = 'z'
	fmt.Printf("%s\n", a)
}

func testSlices3() {
	a := []byte{'a', 'b', 'c'}
	b := append(a[1:2], 'd', 'x')
	b[0] = 'z'
	fmt.Printf("%s\n", a)
}

func testSlices4() {
	a := []byte{'a', 'b', 'c'}
	b := string(a)
	a[0] = 'z'
	fmt.Printf("%s\n", b)
}
