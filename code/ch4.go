package main

import "fmt"

func main() {
	//a1()
	a2()
}

func a1() {
	s := []string{"a", "b", "c"}

	for i, v := range s {
		fmt.Printf("index=%d, value=%s\n", i, v)
	}
	for _, v := range s {
		fmt.Printf("value=%s\n", v)
	}
	for i := range s {
		fmt.Printf("index=%d\n", i)
	}
}

type account struct{ balance float32 }

func a2() {
	accounts := []account{{balance: 100.}, {balance: 200.}, {balance: 300.}}
	for _, a := range accounts {

		a.balance += 1000
	}
	fmt.Println(accounts)
	for i := range accounts {
		accounts[i].balance += 1000
	}
	fmt.Println(accounts)
	for i := 0; i < len(accounts); i++ {
		accounts[i].balance += 1000
	}
	fmt.Println(accounts)
}
