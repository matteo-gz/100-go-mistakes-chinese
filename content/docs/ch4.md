---
title: 04. 控制结构
---

## 30 range下元素copy

{{< hint info >}}
range介绍
{{< /hint >}}
`range` 可作用于 `string`, `array`, `指针数组`, `slice`, `map`, 接收`chan`.

复习下range的写法:

```go
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
```
output:
```shell
index=0, value=a
index=1, value=b
index=2, value=c
value=a
value=b
value=c
index=0
index=1
index=2
```


{{< hint info >}}
range下值类型
{{< /hint >}}
```go
accounts := []account{{balance: 100.}, {balance: 200.}, {balance: 300.}}
for _, a := range accounts {

a.balance += 1000
}
fmt.Println(accounts)
```

> [{100} {200} {300}]

因为我们修改的是拷贝出来的a 原来的slice**没有改变到**

下面才是能改变的

```go
for i := range accounts {
accounts[i].balance += 1000
}
```

或者

```go
for i := 0; i < len(accounts); i++ {
accounts[i].balance += 1000
}

```

那么这种改法呢
`accounts := []*account{}`
作者不是很赞同,在[`#91`][1]时会提及

## 31 range下元素计算

slice

{{< hint info >}}
s在range时被拷贝
{{< /hint >}}


```go
func sliceCopy() {
s := []int{0, 1, 2}
for range s {
s = append(s, 10)
}
fmt.Println(s)
}
```

> [0 1 2 10 10 10]


{{< hint info >}}
陷入死循环,因为s每次迭代都是重新求值
{{< /hint >}}
```go
func sliceNeverStop() {
	s := []int{0, 1, 2}
	for i := 0; i < len(s); i++ {
		s = append(s, 10)
	}
	fmt.Println(s)
}
```
{{< hint info >}}
channel 在range时拷贝值,类似slice
{{< /hint >}}

```go
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
		fmt.Println(v) // 保持ch1的来源,copy in range first time
		// 此时v的来源是ch1的值
		ch = ch2
	}
    fmt.Println("---")
	for v := range ch { // ch 还是被ch2 赋值了
		fmt.Println(v)
	}

}
```
output
```shell
1
11
111
---
2
22
222
```


{{< hint info >}}
array 在range时发生拷贝
{{< /hint >}}
```go
func arrayCopy() {
	arr := [3]int{0, 1, 2}
	for i, v := range arr {
		arr[2] = 10
		if i == 2 {
			fmt.Println(v)
		}
	}
    fmt.Println(arr)
}
```
output
```shell
2
[0 1 10]
```

{{< hint info >}}
array 指针指向源数据
{{< /hint >}}
```go
func arrayCopy() {
	arr := [3]int{0, 1, 2}
	for i, v := range &arr { //此时为指针
		arr[2] = 10
		if i == 2 {
			fmt.Println(v)
		}
	}
    fmt.Println(arr)
}
```
output
```shell
10
[0 1 10]
```
## 32 range下指针元素
range时临时变量指向同个指针
```go

type Customer struct {
	ID      string
	Balance float64
}

type Store struct{ m map[string]*Customer }

func (s *Store) storeCustomers(customers []Customer) {
	for _, customer := range customers {
		// 此时customer 在每次迭代中都是同一指针
		s.m[customer.ID] = &customer
	}

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
```
```shell
&{3 0}
&{3 0}
&{3 0}
```

化解
```go
func (s *Store) storeCustomers(customers []Customer) {
	for _, customer := range customers {
		customer1 := customer
		s.m[customer.ID] = &customer1
	}
}
```
output
```shell
&{1 10}
&{2 -10}
&{3 0}
```

## 33 map迭代陷阱



{{< hint info >}}
map的迭代无序性
{{< /hint >}}


{{< hint info >}}
在同一迭代中操作元素
{{< /hint >}}
```go
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
```
无序的体现,在map range中改变map,每次生成都是无序.

建议用新变量m2来存储赋值结果.

## 34 break工作机制
{{< hint info >}}
break只终结当前循环
{{< /hint >}}


```go
func breakCurr() {
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", i)
		switch i {
		default:
		case 2:
			break
		}
	}
}
```
output
```shell
0 1 2 3 4
```

化解
```go
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
```
output
```shell
0 1 2
```

例子2
```go
func breakFor() {
	ch := make(chan int, 3)
	ctx, canF := context.WithCancel(context.Background())
	go func() {
		time.Sleep(1 * time.Second)
		canF()
	}()
	for {
		select {

		case <-ch:
			fmt.Println("from ch")
		case <-ctx.Done():
			// Do something case <-ctx.Done():
			fmt.Println("wanna break for and exit")
			break
		}
	}
}
```
以上函数无法中终结,只break select

同样修复方式:
```go

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

```
output
```shell
wanna break for and exit
im arrive

```
## 35 在for中使用defer

defer的执行是推入函数的调用栈,在函数**return时才会执行**.

如果在for循环中打开文件,defer关闭文件句柄的行为存在风险

建议其一,将打开文件,defer关闭文件放在一个函数,在for中调用单独函数,保证defer的执行

[1]: ../ch12/#91-cpu缓存