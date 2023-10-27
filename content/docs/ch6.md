---
title: 06. 函数与方法
---

## 42 接收者与指针

没有改变原始数据
```go
func (c customer) add(v float64) { 
	c.balance += v 
}
```

改变原始数据
```go
func (c *customer) add(operation float64) { 
	c.balance += operation 
}
```

## 43 返回值进行命名
```go
getCoordinates(address string) (float32, float32, error)
```
更为清晰
```go
getCoordinates(address string) (lat, lng float32, err error)
```

## 44 返回值非预期
忘记赋值err,导致err是nil
```go
func returnAlwaysNil(ctx context.Context) (err error) {
	if ctx.Err() != nil {
		return err
	}
	return nil
}
```

error被覆盖,本来期望返回"new err"
```go
func returnErrCover(ctx context.Context) (err error) {
	err = errors.New("new err")
	if err := ctx.Err(); err != nil {
		return err // return "context canceled" cover "new err"
	}
	return nil
}
```
## 45 接收者nil

```go

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
// 调用
func show() {
	customer := Customer{Age: 33, Name: "John"}
	if err := customer.Validate(); err != nil {
		fmt.Printf("customer is invalid: %v", err)
	}
}

```
```shell
customer is invalid: <nil>
```

以上本质是`var m *MultiError`

m作为`*MultiError` == nil

但是m却实现了error interface 作为`error` 不是nil

```go
type error interface {
	Error() string
}

```

修复方案

```go
func (c Customer) Validate() error {

    // ...
	
  if m != nil {
    return m
  }

  return nil
}
```


---

简化版理解
```shell
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
```
output:
```shell
err:<nil>
```

涉及知识点为interface的值与类型判断

## 46 文件名作为参数的缺点

文件名作为参数的缺点,不利于单元测试.

解决方案为`io.Reader`作为参数,即接口的优点,抽象代替实际.

## 47 defer arguments与接收者

> 单词parameter与arguments,通常在英文上下文中为形参与实参的区别.

{{< hint info >}}
defer的陷阱
{{< /hint  >}}

```go
func deferFn2(s int) {
	fmt.Println("defer :", s)
}
func deferFn1() {
	var s int
	defer deferFn2(s)
	s = 1
	return
}
```
output:
```shell
defer : 0
```

改法:
```go
func deferFn5() {
	var s int
	defer func() {
		deferFn2(s)
	}()
	s = 1
	return
}
```
output:
````shell
defer : 1
````

{{< hint info >}}
传递通过指针
{{< /hint  >}}

```shell
func deferFn3(s *int) {
	fmt.Println("defer :", *s)
}
func deferFn1() {
	var s int
	defer deferFn3(&s)
	s = 1
	return
}
```
output:
```shell
defer : 1
```

{{< hint info >}}
闭包传递
{{< /hint  >}}
```go
func deferFn4() {
	i := 0
	j := 0
	defer func(i int) {
		fmt.Println(i, j)
	}(i)
	i++
	j++
}
```
output:
```shell
0 1
```

{{< hint info >}}
即刻打印
{{< /hint  >}}

```go
func main2() {
	s := Struct{id: "foo"}
	defer s.print()
	s.id = "bar"
}

type Struct struct {
	id string
}

func (s Struct) print() {
	fmt.Println(s.id)
}
```
output:
```shell
foo
```
{{< hint info >}}
指针类型
{{< /hint  >}}
```go
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
```
output:
````shell
bar
````