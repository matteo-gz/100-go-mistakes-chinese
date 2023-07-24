# 数据类型

## 八进制混淆
即代码相对人类可读性而言
```go
fmt.Println(100 + 010)  //结果 108
fmt.Println(100 + 0o10) //结果 108
fmt.Println(100+010 == 100+0o10) // true
```
八进制表达最好用`0o`前缀

二进制用`0b`前缀

十六进制用`0x`前缀

虚数  使用`i`后缀,例如 `3i`

## 整数溢出
当一个数据超过范围,每次需要增加额外逻辑来判断处理是否溢出。
我个人觉得是一种损耗,应该是你数据类型范围定义错误了。

```go
	var counter int32 = math.MaxInt32
	fmt.Printf("counter=%d\n", counter)
	counter++
	fmt.Printf("counter=%d\n", counter)
```
> counter=2147483647
> 
> counter=-2147483648

## 浮点数理解
具体可以搜索浮点数溢出问题
```go
var n float32 = 1.0001
fmt.Println(n * n) // 应该是1.00020001
```
但是结果是
>  1.0002

## slice底层结构的理解
即len与cap的关系的理解,可以查阅go原理讲解的书籍

### slice 初始化问题
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


if 可读性拉满
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

else 性能拉满
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

### nil与 empty slices 区别
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

### nil处理不当的bug
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
        return operations // 逻辑1 非nil
		return nil // 逻辑2 返回nil
		// 此处应该选择逻辑2,逻辑1有时候人会看走眼
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
### 没有正确copy slice