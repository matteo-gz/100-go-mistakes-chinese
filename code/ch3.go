package main

import "fmt"

func main() {
	fmt.Println(100 + 010)
	fmt.Println(100 + 0o10)
	fmt.Println(100+010 == 100+0o10)
	fmt.Println(0b11)
}
