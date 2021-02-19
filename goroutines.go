package main

import (
	"fmt"
	"runtime"
	"time"
)

func testGoroutines1() {
	var ch chan int
	for i := 0; i < 3; i++ {
		go func(idx int) {
			ch <- (idx + 1) * 2
		}(i)
	}
	fmt.Println("result:", <-ch)
	time.Sleep(2 * time.Second)
}

func testGoroutines2() {
	ch := make(chan string)
	go func() {
		for m := range ch {
			fmt.Println("processed:", m)
		}
	}()
	ch <- "cmd.1"
	ch <- "cmd.2"
}

func testGoroutines3() {
	data := []string{"one", "two", "three"}
	for _, v := range data {
		go func() {
			fmt.Println(v)
		}()
	}
	time.Sleep(3 * time.Second)
}

func testGoroutines4() {
	var num int

	for i := 0; i < 10000; i++ {
		go func() {
			num = i
		}()
	}
	fmt.Printf("NUM is %d", num)
}

func testGoroutines5() {
	dataMap := make(map[string]int)
	for i := 0; i < 10000; i++ {
		go func(d map[string]int, num int) {
			d[fmt.Sprintf("%d", num)] = num
		}(dataMap, i)
	}
	time.Sleep(5 * time.Second)
	fmt.Println(len(dataMap))
}

func testGoroutines6() {
	runtime.GOMAXPROCS(1)

	x := 0
	go func(p *int) {
		for i := 1; i <= 20000000000; i++ {
			*p = i
		}
	}(&x)

	time.Sleep(100 * time.Millisecond)

	fmt.Printf("x = %d.\n", x)
}
