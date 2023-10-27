package main

import (
	"fmt"
	"time"
)

func main() {
	//c2()
	//c1()
	ch3()
}
func c2() {
	i := 0
	go func() {
		i++
	}()
	fmt.Println(i)
}
func c1() {
	i := 0
	ch := make(chan struct{}, 1)
	go func() {
		i = 1
		<-ch
	}()
	ch <- struct{}{}
	fmt.Println(i)
}

var count int

func increment() {
	count++
}

func ch3() {
	go increment()
	go increment()

	time.Sleep(time.Millisecond)

	println(count)
}
