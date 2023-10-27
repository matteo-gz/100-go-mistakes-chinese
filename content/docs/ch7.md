---
title: 07. 错误管理
---

## 48 panic

panic与recover机制

何时panic场景,举例

- net/http 包中存在header code越界panic
- database/sql中 register函数,若driver nil 为panic

## 49 错误修饰

用法

- ` fmt.Errorf("%w",err)` 可显示源错误,用于`errors.Is`判断
- ` fmt.Errorf("%v",err)` 不显示源错误
```go
// 一个原始错误
	underlying := errors.New("underlying error")

	// 使用 %w 报告原始错误
	wrapped := fmt.Errorf("%w", underlying)

	// wrapped 会包含原始错误的信息
	fmt.Println("wrapped:", wrapped) // underlying error

	// 检查 wrapped 是否包含原始错误
	if errors.Is(wrapped, underlying) {
		fmt.Println("%w 报告原始错误")

	}
```
场景

- 遮罩错误,向用户展示人性化提示,并记录错误上下文
- 叠加错误,多个失败堆叠

## 50 错误类型检查
`errors.As` 可用于类型判断
````go
err := DoSomething()

if errors.As(err, &MyError{}); err != nil {
  // err implements MyError
  handleMyError(err)
} else {
  // err does not implement MyError 
  handleGenericError(err)
}
````
## 51 检查错误值不准确
`fmt.Errorf`和`%w`的错误传递的情况下,

用`==`判断错误不可靠,错误有可能被修饰了.

可用`errors.Is`判断.

```go
// err 可能被fmt.Errorf("...%w..." 修饰过
if errors.Is(err, sql.ErrNoRows){
	//...
}
```
## 52 重复处理错误

错误保持精简,多次错误传递时用`%w`组合传递错误

## 53 未处理错误

抑制错误,可能会丢失错误分支
```go
// 函数定义
func notify()error
// 用例

// ...
notify() // 直接调用,抑制错误发生可能
// ...

```

## 54 不处理defer中的错误

建议用日志记录错误
```go
// 函数定义
func close() error

// 用例 直接defer函数,没有处理返回error
defer close()
```
修改后
```go
defer func(){
	err:=close()
	if err!=nil{
		log(err)
    }
}
```