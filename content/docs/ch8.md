---
title: 08. 基础并发
---

## 55 并发与并行
并行即同一时刻,多道任务同时处理.

并发是交替执行,表现在一段时间内完成了多个任务.

## 56 调度的成本

golang协程成本. 文中提及调度原理.

## 57 chan 与 mutex

搭配使用

## 58 data race
多个协程计数器增加时,发生data race,
可用`sync/atomic`,或者`mutex`,`chan`

## 59 工作负载与并发
工作负载因素
- CPU速度
- I/O速度
- 可用内存

上述因素下并发场景的设计

## 60 context

功能

- deadline
  - time.Duration // 从开始持续的时间,比如250 ms
  - time.Time// 时间节点,比如2022-02-02 00:00:00 UTC
- 取消信号
- context values
- context.done
