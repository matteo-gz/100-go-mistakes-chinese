---
title: 5. strings
---

## 36 rune概念

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