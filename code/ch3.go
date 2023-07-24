package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(100 + 010)
	fmt.Println(100 + 0o10)
	fmt.Println(100+010 == 100+0o10)
	fmt.Println(0b11)

	var counter int32 = math.MaxInt32
	fmt.Printf("counter=%d\n", counter)
	counter++
	fmt.Printf("counter=%d\n", counter)

	var n float32 = 1.0001
	fmt.Println(n * n)
}
