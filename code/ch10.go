package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func main() {
	//f2()
	//json1()
	//json2()
	//run1()
	f5()
}
func f5() {
	var i int = 12
	d, err := json.Marshal(i)
	if err != nil {
		fmt.Println(err)
	}
	var m any
	err = json.Unmarshal(d, &m)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%T\n", m)
}
func f2() {
	ch := make(chan int)
	f1(ch)
}
func f1(ch <-chan int) {

	for {
		select {
		case event := <-ch:
			fmt.Println(event)
		case <-time.After(time.Hour):
			log.Println("warning: no messages received")
		}
	}
}

func f3(ch <-chan int) {

	for {
		ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
		select {
		case event := <-ch:
			cancel()
			fmt.Println(event)
		case <-ctx.Done():
			log.Println("warning: no messages received")
		}
	}
}

func f4(ch <-chan int) {
	timeD := time.Hour
	timer := time.NewTimer(timeD)
	for {
		timer.Reset(timeD)
		select {
		case event := <-ch:
			fmt.Println(event)
		case <-timer.C:
			log.Println("warning: no messages received")
		}
	}
}

func json1() {
	event := Event{ID: 1234, Time: time.Now()}

	b, err := json.Marshal(event)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(b))
}

type Event struct {
	ID int
	time.Time
}

func (e Event) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ID   int `json:"id"`
		Time time.Time
	}{
		ID:   e.ID,
		Time: e.Time,
	})
}

type Person struct {
	name string
	age  int
	AB   int
}

type Event1 struct {
	ID     int // 导出
	Person     // 内嵌 匿名
}

func json2() {
	e := Event1{
		ID: 1,
		Person: Person{
			name: "John",
			age:  30,
			AB:   1,
		},
	}

	// 序列化
	data, _ := json.Marshal(e)
	fmt.Println(string(data))
}
func (Person) MarshalJSON() ([]byte, error) {
	return []byte(`2023`), nil
}

type Event3 struct {
	Time time.Time
}

func run1() {
	t := time.Now()
	event1 := Event3{
		Time: t.Truncate(0),
	}

	b, err := json.Marshal(event1)
	if err != nil {
		fmt.Println(err)
		return
	}

	var event2 Event3
	err = json.Unmarshal(b, &event2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(event1 == event2)
	fmt.Println(event1)
	fmt.Println(event2)
	fmt.Println(event1.Time.Equal(event2.Time))
}
