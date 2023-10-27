package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"sync"
	"time"
)

func main1() {
	s := make([]int, 0, 1)

	go func() {
		s1 := append(s, 1)
		fmt.Println(s1)
	}()

	go func() {
		s2 := append(s, 1)
		fmt.Println(s2)
	}()
}

func main() {
	//eg()
	//bka()
	syncCp()
}

func bka() {
	var cond *sync.Cond = sync.NewCond(&sync.Mutex{})
	var balance = 1
	var goal = 10
	//...
	// 读监听
	for i := 0; i < 10; i++ {
		go func(i1 int) {
			cond.L.Lock()
			defer cond.L.Unlock()

			for balance < goal { // loop
				cond.Wait() // 挂起等待唤醒
			}
			fmt.Printf("到达 %d", i1)
		}(i)
	}

	// ...
	// 写
	go func() {

		for {
			func() {
				cond.L.Lock()
				defer cond.L.Unlock()
				balance++
				cond.Broadcast() // 广播
			}()
			//fmt.Println("add")
		}

	}()
	ch := make(chan int)
	<-ch
}
func wg() {
	var wg1 sync.WaitGroup
	wg1.Add(1)
	wg1.Done()
	wg1.Wait()
}

type Counter struct {
	mu       *sync.Mutex
	counters map[string]int
}

func NewCounter() Counter {
	return Counter{
		counters: map[string]int{},
		mu:       &sync.Mutex{},
	}
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
func eg() {
	ctx := context.Background()
	ctx, f := context.WithCancel(ctx)
	defer f()
	g, ctx2 := errgroup.WithContext(ctx)
	g.Go(func() error {
		return biz(ctx2, 1)
	})
	g.Go(func() error {
		return biz(ctx2, 2)
	})
	err := g.Wait()
	fmt.Println(err)
}

func biz(ctx context.Context, v int) error {
	fmt.Println("curr:", v)
	if v == 1 {
		time.Sleep(3 * time.Second)
		return errors.New("sleep and q")
		//return nil
	}
	t1 := time.NewTicker(1 * time.Second)
	for {
		select {
		case ch := <-t1.C:
			fmt.Println("tick")
			fmt.Println(ch)
		case <-ctx.Done():
			fmt.Println("done")
			return nil
		}
	}

}
