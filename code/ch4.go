package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//a41()
	//a2()
	//sliceCopy()
	//sliceNeverStop()
	//chanCopy()
	//arrayCopy()
	//rangeLast()
	//rangeMap()
	//breakCurr()
	breakFor()
}
func breakCurr() {
loopLabel:
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", i)
		switch i {
		default:
		case 2:
			break loopLabel
		}
	}
}
func breakFor() {
	ch := make(chan int, 3)
	ctx, canF := context.WithCancel(context.Background())
	go func() {
		time.Sleep(1 * time.Second)
		canF()
	}()
loopLabel:
	for {
		select {

		case <-ch:
			fmt.Println("from ch")
			break
		case <-ctx.Done():
			// Do something case <-ctx.Done():
			fmt.Println("wanna break for and exit")
			break loopLabel
		}
	}
	fmt.Println("im arrive")
}

type Customer struct {
	ID      string
	Balance float64
}

type Store struct{ m map[string]*Customer }

func (s *Store) storeCustomers(customers []Customer) {
	for _, customer := range customers {
		customer1 := customer
		s.m[customer.ID] = &customer1
	}
}
func rangeMap() {
	m := map[int]bool{
		0: true, 1: false, 2: true,
	}

	for k, v := range m {
		if v {
			m[10+k] = true
		}
	}
	fmt.Println(m)
}
func rangeLast() {
	s := Store{
		m: map[string]*Customer{},
	}
	s.storeCustomers([]Customer{{ID: "1", Balance: 10}, {ID: "2", Balance: -10}, {ID: "3", Balance: 0}})
	for _, v := range s.m {
		fmt.Println(v)
	}
}
func sliceCopy() {
	s := []int{0, 1, 2}
	for range s {
		s = append(s, 10)
	}
	fmt.Println(s)
}
func sliceNeverStop() {
	s := []int{0, 1, 2}
	for i := 0; i < len(s); i++ {
		s = append(s, 10)
	}
	fmt.Println(s)
}
func chanCopy() {
	ch1 := make(chan int, 3)
	ch2 := make(chan int, 3)
	go func() {
		ch1 <- 1
		ch1 <- 11
		ch1 <- 111
		close(ch1)
	}()
	go func() {
		ch2 <- 2
		ch2 <- 22
		ch2 <- 222
		close(ch2)
	}()
	ch := ch1
	for v := range ch {
		fmt.Println(v)
		ch = ch2
	}
	fmt.Println("---")
	for v := range ch {
		fmt.Println(v)
	}

}
func arrayCopy() {
	arr := [3]int{0, 1, 2}
	for i, v := range &arr {
		arr[2] = 10
		if i == 2 {
			fmt.Println(v)
		}
	}
	fmt.Println(arr)
}
func a41() {
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
