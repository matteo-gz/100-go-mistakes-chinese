---
title: 05. strings
---

## 36 string相关rune概念

- charset
- encoding,golang使用UTF-8

```go
s := "hello"

fmt.Println(len(s)) // 5

s := "汉" 
fmt.Println(len(s)) // 3

s := string([]byte{0xE6, 0xB1, 0x89}) 
fmt.Printf("%s\n", s)// 汉
```

rune 即 unicode.

len 返回的byte count.

## 37 strings迭代
```go
s := "hêllo"

for i := range s {
	fmt.Printf("position %d: %c\n", i, s[i])
} 
fmt.Printf("len=%d\n", len(s))
```
output:
```shell
position 0: h
position 1: Ã
position 3: l
position 4: l
position 5: o
len=6
```

---

改用rune
```go
s := "hêllo"
runes := []rune(s)
for i, r := range runes {
fmt.Printf("position %d: %c\n", i, r)
}
fmt.Printf("len=%d\n", len(runes))
```
output:
```shell
position 0: h
position 1: ê
position 2: l
position 3: l
position 4: o
len=5
```

string默认在byte数组中的迭代,底层用的UTF-8识别.

len的计算可以使用`utf8.RuneCountInString(s)`获取真实rune数.

`[]rune(s)`的转换并非魔法,使用时产生了运行时计算开销.

## 38 string相关trim函数

```go
	fmt.Println(strings.TrimRight("123oxo", "xo"))  // 123
	fmt.Println(strings.TrimSuffix("123oxo", "xo")) //123o

	fmt.Println(strings.TrimLeft("oxo123", "ox"))   // 123
	fmt.Println(strings.TrimPrefix("oxo123", "ox")) /// o123

	fmt.Println(strings.Trim("oxo123oxo", "ox")) // 123
```
具体的函数定义处有详细说明

## 39 string接连

对于少量字符拼接,保持易读性.
使用`fmt` `+=` 拼接都可以

对于大量字符,使用builder函数,并且设置grow可有良好的优化.

写多个拼接函数版本,然后压测即可.上述的只是结论.
## 40 string转换

很多接口的操作是[]byte操作,改为string操作存在转换消耗.

## 41 substrings与内存泄漏
占据同个内存,可见[`#26`][1]
```go
uuid := log[:36] 
s.store(uuid)
```

改为让slice脱离释放
```go
uuid := strings.Clone(log[:36])
```

[1]: ../ch3/#26-slice-与内存泄漏