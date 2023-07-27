package main

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

func main() {
	//fmt.Println(100 + 010)
	//fmt.Println(100 + 0o10)
	//fmt.Println(100+010 == 100+0o10)
	//fmt.Println(0b11)
	//
	//var counter int32 = math.MaxInt32
	//fmt.Printf("counter=%d\n", counter)
	//counter++
	//fmt.Printf("counter=%d\n", counter)
	//
	//var n float32 = 1.0001
	//fmt.Println(n * n)
	//
	//sl()
	//sl1()
	//sl2()
	//sl3()
	//sl4()
	//sl5()

	//consumeMessages()
	//a1()
	//mapGc()
	eq()
}
func eq() {
	fmt.Println(struct{ id string }{"1"} == struct{ id string }{"1"})
	type customer struct {
		id         string
		operations []float64
	}
	var c1 any = 1
	var c12 any = 1
	fmt.Println(c1 == c12)
	var c11 any = customer{id: "x", operations: []float64{1.}}
	var c13 any = customer{id: "x", operations: []float64{1}}
	fmt.Println(reflect.DeepEqual(c11, c13))
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

func consumeMessages() {
	i := 1000
	var m runtime.MemStats
	for {

		runtime.ReadMemStats(&m)

		fmt.Println("Allocated memory (bytes):", m.Alloc)
		// getMessageType 2199560
		// getMessageType 1199236608
		if i == 0 {
			break
		}
		i--
		msg := receiveMessage()
		storeMessageType(getMessageType2(msg))
	}
	fmt.Println("ok")
}
func receiveMessage() []byte {
	s := "1"
	b := strings.Repeat(s, 100*10000)
	return []byte(b)
}

var a [][]byte

func storeMessageType(b []byte) {
	a = append(a, b)
}
func getMessageType(msg []byte) []byte {
	return msg[:5]
}
func getMessageType2(msg []byte) []byte {
	msgType := make([]byte, 5)
	copy(msgType, msg)
	return msgType
}

type Foo struct{ v []byte }

func printAlloc() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d KB\n", m.Alloc/1024)
}
func a1() {
	foos := make([]Foo, 1_000)
	printAlloc()
	for i := 0; i < len(foos); i++ {
		foos[i] = Foo{
			v: make([]byte, 1024*1024),
		}
	}
	printAlloc()
	two := keepFirstTwoElementsOnly3(foos)
	runtime.GC()
	printAlloc()
	runtime.KeepAlive(two)
}
func keepFirstTwoElementsOnly(foos []Foo) []Foo {
	return foos[:2]
}
func keepFirstTwoElementsOnly2(foos []Foo) []Foo {
	res := make([]Foo, 2)
	copy(res, foos)
	return res
}
func keepFirstTwoElementsOnly3(foos []Foo) []Foo {
	for i := 2; i < len(foos); i++ {
		foos[i].v = nil
	}
	return foos[:2]
}

func mapGc() {
	n := 1_000_000
	m := make(map[int][128]byte)
	printAlloc()
	for i := 0; i < n; i++ {
		m[i] = randBytes()
	}
	printAlloc()
	for i := 0; i < n; i++ {
		delete(m, i)
	}
	runtime.GC()
	printAlloc()
	runtime.KeepAlive(m)
}
func randBytes() [128]byte {
	return [128]byte{}
}
