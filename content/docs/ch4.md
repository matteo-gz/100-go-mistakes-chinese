---
title: 4. 控制结构
---
## 30 元素在循环中复制

`range` 可作用于 string array 指针数组 slice map 接收chan

复习下

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

### 开始陷阱挑战

```go
accounts := []account{{balance: 100.}, {balance: 200.}, {balance: 300.}}
 for _, a := range accounts {

  a.balance += 1000
 }
 fmt.Println(accounts)
```

> [{100} {200} {300}]

因为我们修改的是拷贝出来的a 原来的slice没有改变到

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
作者不是很赞同,在`#91`时会提及

## 31 忽略如何在range循环中计算参数


