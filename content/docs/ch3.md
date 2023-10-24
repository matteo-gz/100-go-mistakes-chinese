---
title: 3. 数据类型
---


## 17 八进制混淆

即代码相对人类可读性而言.

```go
fmt.Println(100 + 010)  //结果 108
fmt.Println(100 + 0o10) //结果 108
fmt.Println(100+010 == 100+0o10) // true
```

八进制表达最好用`0o`前缀.

二进制用`0b`前缀.

十六进制用`0x`前缀.

虚数  使用`i`后缀,例如 `3i`.

## 18 整数溢出

当一个数据超过范围,每次需要增加额外逻辑来判断处理是否溢出.
个人觉得是一种损耗,属于数据范围应用不当.

```go
 var counter int32 = math.MaxInt32
 fmt.Printf("counter=%d\n", counter)
 counter++
 fmt.Printf("counter=%d\n", counter)
```

> counter=2147483647
>
> counter=-2147483648

## 19 浮点数理解

具体可以搜索浮点数溢出问题.

```go
var n float32 = 1.0001
fmt.Println(n * n) // 应该是1.00020001
```

但是结果是
> 1.0002

## 20 slice底层结构的理解


即len与cap的关系的理解.

## 21 slice 初始化问题

```go
func convert(foos []Foo) []Bar {
 // 第一种低效的
 bars := make([]Bar, 0) 
 // 第二种高效的
 n := len(foos) 
 bars := make([]Bar, 0, n)
 // 因为不知道slice里面的元素数量,如果slice过大则需要多次扩容
 for _, foo := range foos {
  bars = append(bars, fooToBar(foo))

 }
 return bars
}
```

### 关于append的处理

偏向可读性做法:

```go
func collectAllUserKeys(cmp Compare, tombstones []tombstoneWithLevel) [][]byte {
    keys := make([][]byte, 0, len(tombstones)*2)
    for _, t := range tombstones {
        keys = append(keys, t.Start.UserKey)
        keys = append(keys, t.End)
    }
    // ...
}

```

偏向性能做法:

```go
func collectAllUserKeys(cmp Compare, tombstones []tombstoneWithLevel) [][]byte {
    keys := make([][]byte, len(tombstones)*2)
    for i, t := range tombstones {
        keys[i*2] = t.Start.UserKey
        keys[i*2+1] = t.End
    }
    // ...
}

```

2种写法见仁见智,如果调用次数多,成为性能瓶颈情况下.还是为了团队沟通,让代码可读性变高.

## 22 nil与 empty slices 区别

- len==0 就是empty slices
- slice==nil 就是 nil slice

```go
var s []string  // empty 且nil

s = []string(nil) // empty 且nil

s = []string{} // 只是empty

s = make([]string, 0) // 只是empty
```

以上区别就是nil是没有内存分配的

### json encode影响

```go
var s1 []float32  // nil 
customer1 := customer{ ID: "foo", Operations: s1, }
```

> {"ID":"foo","Operations":null}

第2种

```go
s2 := make([]float32, 0) // 非nil
customer2 := customer{ ID: "bar", Operations: s2, }
```

> {"ID":"bar","Operations":[]}

## 23 nil处理不当的bug

```go
func handleOperations(id string) { 
 operations := getOperations(id) 
 if operations != nil {
  // 如果有操作就是要去处理
  handle(operations) 
 } 
}
func getOperations(id string) []float32 { 
 operations := make([]float32, 0) // 此处已经非nil了
    if id == "" {
        return operations // 逻辑1 误以为是nil,其实此处返回了非nil
  return nil // 逻辑2 返回nil
  // 此处应该选择逻辑2处理才是正确的
    }
    operations=append(operations,1.0)
    return operations // 非nil
}
```

另一种方案

```go
func handleOperations(id string) { 
 operations := getOperations(id) 
 //if operations != nil {
 // 改为判断长度
 if len(operations) != 0 {
  // 如果有操作就是要去处理
  handle(operations) 
 } 
}
```

## 24 没有正确copy slice

下面例子

错误的

```go
    src := []int{0, 1, 2}
 var dst []int
 copy(dst, src)
    fmt.Println("dst:", dst,len(dst),cap(dst))
```

output:
> dst: [] 0 0

---

正确的

```go
    src := []int{0, 1, 2}
dst := make([]int, len(src))
 copy(dst, src)
    fmt.Println("dst:", dst)
```

output:
> dst: [0 1 2]

---

这种语法也能拷贝的

```go
    src := []int{0, 1, 2}
 dst := append([]int(nil), src...)
 fmt.Println("dst:", dst)
```

output:
> dst: [0 1 2]

## 25 append的注意点

```go
s1 := []int{1, 2, 3}

 s2 := s1[1:2]

 s3 := append(s2, 10)
 fmt.Println(s1, s2, s3)
```

结果会是怎样呢?

或许你会猜测是:
`//  [1,2,3] [2]  [2,10]`

但真实的output:

> [1 2 10] [2] [2 10]

因为他们指向同个底层array,当append发生,len没有超出cap时,s1[2]被修改到了.

**类似同样的现象**

```go
func main() {

s := []int{1, 2, 3}

// s[:2]// [1,2] but cap is 3
f(s[:2])

fmt.Println(s) // [1 2 10]

}
func f(s []int) { 
 _ = append(s, 10) 
}
```

那么如何保护slice,不受上下文被改动到呢?

- 通过copy函数,使用新变量
- 利用s[low:high:max] 这种的表达式,cap==max-low

```go
f(s[:2:2]) // 即s[0:2:2] cap==2-0 ->2
// 调用时则不会改动原有的s
```

## 26 slice与内存泄漏

例子 我们调用`consumeMessages`

```go
func consumeMessages() {
 i := 1000
 var m runtime.MemStats
 for {

  runtime.ReadMemStats(&m)

  fmt.Println("Allocated memory (bytes):", m.Alloc)
  // getMessageType 2199560
  // getMessageType2 1199236608
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
 copy(msgType, msg) // 新的变量
 return msgType
}
```

`getMessageType`与`getMessageType2`相差内存占有不一样
> getMessageType 2199560
>
> getMessageType2 1199236608
`getMessageType2`方法是有效的.

那么如果`getMessageType`下面这种处理有效吗?

```go
func getMessageType(msg []byte) []byte {
 return msg[:5:5] // 之前的写法是 msg[:5]
}
```

答案是改成`msg[:5:5]`也是无效处理

---

**slice与指针**

例子
我们调用`a1`

第一种 无法垃圾回收

```go

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
 two := keepFirstTwoElementsOnly(foos)
 runtime.GC()
 printAlloc()
 runtime.KeepAlive(two) // 此时 foos 底层数组被强行留着
}
func keepFirstTwoElementsOnly(foos []Foo) []Foo {
 return foos[:2] 
}
```

> 104 KB
>
> 1024108 KB
>
> 1024109 KB

---

第二种 可以回收

```go
func keepFirstTwoElementsOnly2(foos []Foo) []Foo {
 res := make([]Foo, 2)
 // 消耗成本取决于cap(res)的大小
 copy(res, foos)
 return res
}
```

> 104 KB
>
> 1024110 KB
>
> 2159 KB

---

第三种 也可以垃圾回收

```go
func keepFirstTwoElementsOnly3(foos []Foo) []Foo {
 // 消耗成本取决于len(foos)的大小
 for i := 2; i < len(foos); i++ {
  foos[i].v = nil
 }
 return foos[:2]
}

```

> 104 KB
>
> 1024107 KB
>
> 2157 KB

关于第二种和第三种做法,哪种好取决于你的场景以及做的基准测试.

## 27 map初始化

```go
// 有初始化容量的
m := make(map[string]int, 1_000_000)
```

```go
BenchmarkMapWithoutSize-4  6   227413490 ns/op
BenchmarkMapWithSize-4    13    91174193 ns/op
```

作者压测后得出有初始化数量更高效.

文中解释了这一现象原因:一个合理数量的初始化,不用动态创建bucket以及重新平衡bucket,从而高效.

## 28 map与内存泄漏

例子

```go
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
```

> 104 KB
>
> 472504 KB
>
> 300441 KB

结论是垃圾回收了,但是没有想象中的回收的多,此处涉及map的底层数据结构.

存在的bucket没变化,只是里面的slots变成了0,map只能不断的增长,拥有更多的bucket,而不会缩小

探讨:
一场活动导致的流量高峰,高峰过后因为map的占用内存导致服务器的内存处于高位,而回收不太理想

- 每隔一小时复制到新map,丢弃旧map
- 优化数据类型 `map[int][128]byte`变成`map[int]*[128]byte`,改为指针能节省部分内存

## 29 判断对等关系

`==` 不能用于slice map

`reflect.DeepEqual()` 可以做到,但是性能很一般
