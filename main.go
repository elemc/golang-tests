package main

import "fmt"

func main() {
	fmt.Printf("\n")
	fmt.Printf("iota tests...\n")
	testIOTA1()
	testIOTA2()

	fmt.Printf("\n")

	fmt.Printf("Arrays and slices tests...\n")
	testSlices1()
	testSlices2()
	testSlices3()
	testSlices4()

	fmt.Printf("\n")

	fmt.Printf("goroutines tests...\n")
	testGoroutines1()
	testGoroutines2()
	testGoroutines3()
	testGoroutines4()
	testGoroutines5()
	testGoroutines6()

	fmt.Printf("\n")

	fmt.Printf("data tests...\n")
	testData1()
	testData2()
	fmt.Printf("\n")
}
