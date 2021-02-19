package main

import "fmt"

func testIOTA1() {
	const (
		p1 = iota
		q1 = iota
		r1 = iota
	)

	fmt.Println(p1, q1, r1)
}

func testIOTA2() {
	const (
		p2, q2, r2 = iota, iota, iota
	)

	fmt.Println(p2, q2, r2)
}
