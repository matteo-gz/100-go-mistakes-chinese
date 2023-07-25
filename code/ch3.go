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

	sl()
	sl1()
	sl2()
	sl3()
	sl4()
	sl5()
}
func sl() {
	src := []int{0, 1, 2}
	var dst []int
	copy(dst, src)
	fmt.Println("dst:", dst, len(dst), cap(dst))
}
func sl1() {
	src := []int{0, 1, 2}
	dst := make([]int, len(src))
	copy(dst, src)
	fmt.Println("dst:", dst)
}
func sl2() {
	src := []int{0, 1, 2}
	dst := append([]int(nil), src...)
	fmt.Println("dst:", dst)
}
func sl3() {
	s1 := []int{1, 2, 3}

	s2 := s1[1:2]

	s3 := append(s2, 10)
	fmt.Println(s1, s2, s3)
	fmt.Println(len(s1), len(s2), len(s3))
	fmt.Println(cap(s1), cap(s2), cap(s3))
	//  123 2  2,10
	// 1 2 10
}

func sl4() {
	s := []int{1, 2, 3}
	fmt.Println(s[:2])
	fmt.Println(cap(s[:2]))
	f(s[:2])

	fmt.Println(s) // [1 2 10]
}
func f(s []int) {
	_ = append(s, 10)
}
func sl5() {
	s := []int{1, 2, 3}
	fmt.Println(s[:2:2])
	fmt.Println(cap(s[:2:2]))
	f(s[:2:2])

	fmt.Println(s) // [1 2 10]
}
