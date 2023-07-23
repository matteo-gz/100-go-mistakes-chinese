package main

import (
	"fmt"
	"sync"
)

var va = func() int {
	fmt.Println("var")
	return 1
}()

func init() {
	fmt.Println("init")
}
func init() {
	fmt.Println("init2")
}
func main() {
	fmt.Println("main")

}

var va1 = func() int {
	fmt.Println("var1")
	return 1
}()

type a struct {
	sync.Mutex
	mu sync.Mutex
}

func useA() {
	a1 := new(a)
	a1.Lock()
	a1.mu.Lock()
}
