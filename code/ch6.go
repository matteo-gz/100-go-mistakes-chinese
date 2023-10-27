package main

import (
	"context"
	"errors"
	"fmt"
	"strings"
)

func main() {
	//l := loc{}
	//c, f1 := context.WithCancel(context.Background())
	//f1()
	//fmt.Println(l.getCoordinates(c, ""))
	//show()
	//T1()
	//simpleErr()
	//deferFn5()
	main2()
}
func main2() {
	s := &Struct2{id: "foo"}
	defer s.print()
	s.id = "bar"
}

type Struct2 struct {
	id string
}

func (s *Struct2) print() {
	fmt.Println(s.id)
}
func deferFn4() {
	i := 0
	j := 0
	defer func(i int) {
		fmt.Println(i, j)
	}(i)
	i++
	j++
}
func deferFn2(s int) {
	fmt.Println("defer :", s)
}
func deferFn3(s *int) {
	fmt.Println("defer :", *s)
}
func deferFn1() {
	var s int
	defer deferFn3(&s)
	s = 1
	return
}
func deferFn5() {
	var s int
	defer func() {
		deferFn2(s)
	}()
	s = 1
	return
}

type mErr struct {
	errs []string
}

func (m *mErr) Error() string {
	return strings.Join(m.errs, ";")
}
func simpleErr() {
	var m *mErr
	var e2 error
	e2 = m
	if err := e2; err != nil {
		fmt.Printf("err:%v", err)
	}
}

type loc struct {
}

func (l loc) validateAddress(string2 string) bool {
	return true
}
func returnAlwaysNil(ctx context.Context) (err error) {
	if ctx.Err() != nil {
		return err
	}
	return nil
}
func returnErrCover(ctx context.Context) (err error) {
	err = errors.New("new err")
	if err := ctx.Err(); err != nil {
		return err // return "context canceled" cover "new err"
	}
	return nil
}
func (l loc) getCoordinates(ctx context.Context, address string) (lat, lng float32, err error) {
	isValid := l.validateAddress(address)
	if !isValid {
		return 0, 0, errors.New("invalid address")
	}
	err = errors.New("new err")
	if err := ctx.Err(); err != nil {
		//fmt.Println(ctx.Err())
		return 0, 0, err
	}
	//Checks whether the context was canceled or the deadline has passed

	// Get and return coordinates
	return
}

type MultiError struct {
	errs []string
}

func (m *MultiError) Add(err error) {

	m.errs = append(m.errs, err.Error())
}
func (m *MultiError) Error() string {
	return strings.Join(m.errs, ";")
}

type Customer struct {
	Age  int
	Name string
}

func (c Customer) Validate() error {

	var m *MultiError
	if c.Age < 0 {
		m = &MultiError{}
		m.Add(errors.New("age is negative"))
	}
	if c.Name == "" {
		if m == nil {
			m = &MultiError{}
		}
		m.Add(errors.New("name is nil"))
	}
	return m

}
func show() {
	customer := Customer{Age: 33, Name: "John"}
	if err := customer.Validate(); err != nil {
		fmt.Printf("customer is invalid: %v", err)
	}
}

type Foo struct{}

func (foo *Foo) Bar() string {
	return "bar"
}
func T1() {
	var foo *Foo
	fmt.Println(foo)
	fmt.Println(foo.Bar())
}
