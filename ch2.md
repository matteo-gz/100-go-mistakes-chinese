# 第2章:代码与项目组织

## 变量重名
```go
package main

import "fmt"

func main() {
	// 外部作用域变量
	x := 10

	// 这是一个新的代码块，内部作用域
	{
		// 在内部作用域中重名了一个变量x
		x := 5
		fmt.Println("内部作用域中的x:", x) // 输出：内部作用域中的x: 5
	}

	// 注意：外部作用域中的变量x没有被修改，因为它被内部作用域中的x所掩盖
	fmt.Println("外部作用域中的x:", x) // 输出：外部作用域中的x: 10
}

```
output:
>
> 内部作用域中的x: 5
> 
> 外部作用域中的x: 10
>

解决办法

这年头,IDE goland 已经能识别这个风险了,新作用域内的x变量会变成绿色字体的
![绿色变量](./code/ch2.png)

用一个临时变量tmpX来过渡,防止2个地方的`x`变量指向的混淆
```go
x:=10
{
	tmpX:=5
	x=tmpX
}
```

要么不用`:=`
```go
var x int
x=10
{
	x=5
}
```

## 代码嵌套
即`卫语句`概念,不懂的可以wiki下
```
if（it == 活的）{
 

    if（it == 人）{
 

        if（it ！= 女人）{
 

            return 不喜欢；

        } else {
 

            return 喜欢；

        }

    } else {
 

        return 不喜欢；

    }

} else {
 

    return 不喜欢；

}
```
改成如下
```
if （it ！= 活的）{return 不喜欢}

if（it！=人）{return 不喜欢}

if（it！=女人）{return 不喜欢}

if（it == 女人 && it == 活的） {return 喜欢}

if（其他任何情况）{return 不喜欢}
```

## init 函数使用场景
例子1 

init与其他函数执行顺序
```go
package main

import "fmt"
var va = func() int {
	fmt.Println("var")
	return 1
}()

func init() {
	fmt.Println("init")
}
func main() {
	fmt.Println("main")
}
```
output:
>
> var
> 
> init
> 
> main


例子2
```go

func main(){
	redis.xx
}
```
先载入redis的init,然后是main的init

例子3

import多个包时,根据文件字母顺序载入init,比如b.go,a.go.
先执行a.go的init


例子4

init函数可以多处定义

```go
func init() {
fmt.Println("init")
}
func init() {
fmt.Println("init2")
}
```
先执行第一个init.




- 初始化其他包的init

```go
import (
_ "foo"
)

```
上面例子中foo包的init被执行了

- init函数无法被其他函数调用


### 使用场景
#### 反例
数据库连接例子写在init函数里可能是不合适的
``` 
// 伪代码
var db *DB

fun init(){
    db，err=xx()
    if err!=nil{
        painc 
    }
}
```
- init函数没有返回值,只能panic去中断,而数据库返回的error,只能暴力处理
- 考虑到单元测试场景,init是首次被加载的,包里的函数不是每个测试时都需要数据库连接这个依赖项
- 使得变量全局化了,过于暴露了数据库这个变量

#### 正例
```go
func init(){
http.HandleFunc("/blog/", redirect)
}
```
- error的handle还是正常的
- 单元测试没有变复杂

## 过渡封装函数来设置变量值

或许你见过这种函数对
```go
func setAge(){ //setter
	
}
func Age(){ // getter
	
}
```
他们的使用好处是
- 统一管理
- 可以在函数内部对变量作出规则限制
- 很方便植入断点

如果前期简单需求,请不要对变量过度封装

## 接口污染

接口越大,抽象越弱

定义接口规则

- 共同行为
- 解耦合
- 约束行为

不要去设计接口,而是发现他们.

接口也是有cpu消耗成本的.

接口的定义尽量在消费端,而不是在生产端

函数返回对象尽量不要用interface,否则别人使用时,需要去看你的代码

入参可以接受接口,宽进严出.

## any类型使用
any使得静态语言变得和动态语言一样,不确定里面的信息.
除非在json encode和decode场景这种,尽量减少用any,因为他代表着信息表征减少.

## 泛型使用

常见:
- 比如slice里面的元素合并

作者说了句,go使用上已经很久没用泛型了(因为之前没引入泛型).泛型使用,见仁见智。

## 内嵌类型
```go
type a struct {
	sync.Mutex
}

func useA() {
	a1 := new(a)
	a1.Lock()
}
```
a1.Lock 对于使用者过于迷惑
```go
type a struct {
	mu sync.Mutex
}

func useA() {
	a1 := new(a)
	a1.mu.Lock()
}
```
a1.mu.Lock 这样对于使用者就不会感到迷惑了。



对于`a.c` 还是`a.b.c`,哪种更好,要具体情况具体分析.

## 介绍了func option 模式优点

帮助你联想的代码
```go

func WithPort(port int) Option {
  return func(s *Server) {
    s.port = port
  }
}
 
func WithTimeout(timeout time.Duration) Option {
  return func(s *Server) {
    s.timeout = timeout
  }
}

```
如果你不知道的化,具体可以搜索下`functional options pattern golang`

## 介绍了go标准布局
作者安利了`https://github.com/golang-standards/project-layout`

## 包名定义
减少这种`common`包的定义,这里涉及了代码规范问题,见仁见智.

## 包名冲突
使用别名解决(有点无语,这也能说道说道的吗...)

## 代码文档
写代码文档.
这个属于代码规范问题了.

## 介绍了linter
linter,静态代码检测.
安利了集成的linter,项目`https://github.com/golangci/golangci-lint`


## 私货
这里安利下uber公司的代码规范 `https://github.com/xxjwxc/uber_go_guide_cn`