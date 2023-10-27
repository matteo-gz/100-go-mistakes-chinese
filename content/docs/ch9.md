---
title: 09. 实践并发
---

## 61 context传递错误
若一个context变量无法应对多个场景

- 新建context.background
- 建造新context,只传递context values

## 62 协程何时终止

goroutine 泄漏

记得关闭的资源
- HTTP
- 数据库连接
- 打开文件
- socket

资源的关闭机制

在main函数执行时,通过context的传递来控制协程,或者defer关闭(作者推荐这个,因为defer保证执行).

```go

w:=newWatcher()
defer w.close()

// 在对应的close 触发终止
func (watcher) close (){
	// 写上资源的回收
}
```

## 63 协程的时序
在for循环中不断地创建协程,他们执行的顺序是没有先后的.

## 64 select与channels
select与channels搭配保证时序
```go
select {
 case v:=<-ch:
}
```

## 65 信号与channels
`chan struct{}` 承担信号的变化,减少内存占用和符号意义代表.
## 66 nil channels
- 利用nil chan阻塞
- 合并chan消息时,用nil chan避免CPU空转

## 67 channel size
size为0,易于推理逻辑,同步模式

size为40,CPU与内存平衡

size过大要思考crash时场景以及内存占用.

## 68 string格式

 

1个协程context更新值;1个协程fmt打印context造成数据竞争,案例:https://github.com/etcd-io/etcd/pull/7816


fmt 打印结构体时会调用String()接口

## 69 data race与append
存在多个goroutine append slice操作时导致data race.

因为slice可能指向同个底层数组,在数据修改时就存在数据竞争.

## 70 mutex与map
在考虑map的data race,确定mutex的边界,在保证线程安全情况下,提高并发程度,减少锁占用时间.

## 71 waitGroup误用
```go
go func(){
	wg.Add(1) //wg的数量存在情况:即由于协程有几率执行比wg.wait慢,从而挂不上.
	wg.Done()
}
wg.Wait()
```
这种情况下存在bug

```go
wg.Add(1) // wg的数量已经挂上
go func(){
    wg.Done()
}
wg.Wait()
```
以上保证了wg的数量有被挂上,而非goroutine数量过多时,得不到执行.

## 72 sync.Cond
监听balance是否到达goal,mutex的实现例子
```go
var mu sync.RWMutex


//...
// 读锁 监听
mu.RLock()
defer mu.RUnlock()

for balance< goal{ // 陷入loop CPU损耗
	mu.RUnlock()
	mu.RLock()
}

fmt.Printf("到达")
//...
// 写锁 操作数据
mu.Lock()
defer mu.Unlock()

balance++

```
上述场景,每次去窥探是否到达goal时,一直陷入循环.

---

利用chan改造
```go
ch:=make (chan int)
//...
//读取 监听
for balance:=range ch{
	if balance>= goal{
	    return	
    }   
}
// ...
// 写入
balance++
ch<-balance
```
当写入有变化时,对比第一种方案有所优化.

但是只能设置一个读取函数,若存在同时监听2个目标时,比如:

balance与 `goal=10`, `goal=11`比较时chan只能被读取一次,下次无法读取了.

---

cond 广播机制
```go
var cond *sync.Cond=sync.NewCond(&sync.Mutex{})
//...
// 读监听
cond.L.Lock()
defer cond.L.Unlock()

for balance<goal { // loop
	cond.Wait() // 挂起等待唤醒
}

fmt.Printf("到达")
// ...
// 写
cond.L.Lock()
defer cond.L.Unlock()

balance++

cond.Broadcast() // 广播
```
利用广播,可满足**多个goal设置值监听**
## 73 errgroup
多goroutine协助,errgroup妙用
```go
//...
g,ctx:=errgroup.WithContext(ctx)
for _,v:= range data{
	g.Go(
		func()error{
		    err:=xx(ctx,v) // xx 为业务函数
			return err
        }
	)
}
err:=g.Wait()
```

引入errgroup包,使用WithContext来创建一个errgroup,并传入上下文ctx。

使用errgroup的Go方法,为每个data循环起一个goroutine调用业务函数xx。这里xx函数必须接收ctx作为参数,因为需要**读取ctx中的取消信号**。

errgroup的Go方法会将每个任务以goroutine的形式运行,并将错误返回替换为errgroup中管理的错误。

调用errgroup的Wait方法,会**阻塞等待**所有任务完成。如果其中有任务返回错误,则Wait最终会返回**第一个错误**。

可以在ctx中设置取消信号,此时所有goroutine都会收到取消消息退出。

## 73 复制sync

**sync包永远不该被复制类型**

- sync.Cond
- sync.Map
- sync.Mutex
- sync.RWMutex
- sync.Once
- sync.Pool
- sync.WaitGroup

案例:
```go

type Counter struct {
	mu       sync.Mutex
	counters map[string]int
}

func NewCounter() Counter {
	return Counter{counters: map[string]int{}}
}

func (c Counter) Increment(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}
func syncCp() {
	counter := NewCounter()
	for i := 0; i < 100; i++ {
		go func() {
			counter.Increment("foo")
		}()
		go func() {
			counter.Increment("bar")
		}()
	}
}
```
`syncCp`调用会
```shell
fatal error: concurrent map writes
```

修复方法1:

IDE提示为
> 'Increment' passes a lock by the value: type 'Counter' contains 'sync.Mutex' which is 'sync.Locker'


修复后
```go
func (c *Counter) Increment(name string) //修改为指针接收者
```

修复方法2:
```go
type Counter struct {
	mu       *sync.Mutex  // 指针类型
	counters map[string]int
}
func NewCounter() Counter {
    return Counter{
        counters: map[string]int{},
        mu:       &sync.Mutex{}, // 指向同个,若此处为nil 也会panic
    }
}
```