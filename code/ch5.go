package main

import (
	"fmt"
	"strings"
)

func main() {
	//rune1()
	//iter()
	//iter2()
	trim1()
}
func rune1() {
	s := "hello"

	fmt.Println(len(s)) // 5

	s = "汉"
	fmt.Println(len(s)) // 3

	s = string([]byte{0xE6, 0xB1, 0x89})
	fmt.Printf("%s\n", s)
}

func iter() {
	s := "hêllo"

	for i := range s {
		fmt.Printf("position %d: %c\n", i, s[i])
	}
	fmt.Printf("len=%d\n", len(s))
}

func iter2() {
	s := "hêllo"
	runes := []rune(s)
	for i, r := range runes {
		fmt.Printf("position %d: %c\n", i, r)
	}
	fmt.Printf("len=%d\n", len(runes))
}
func trim1() {
	fmt.Println(strings.TrimRight("123oxo", "xo"))  // 123
	fmt.Println(strings.TrimSuffix("123oxo", "xo")) //123o

	fmt.Println(strings.TrimLeft("oxo123", "ox"))   // 123
	fmt.Println(strings.TrimPrefix("oxo123", "ox")) /// o123

	fmt.Println(strings.Trim("oxo123oxo", "ox")) // 123
}
