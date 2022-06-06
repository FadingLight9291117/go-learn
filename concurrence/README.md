# Goroutine 协程

## goroutine

[示例代码](./goroutine/goroutine.go)

runtime 库

- `runtime.GOMAXPROCS(0)` 设置和查询可以并发执行的 goroutine 数(参数 0 为查询)，默认值为核心数

- `runtime.NumGroutine()` 返回当前 goroutine 数目，main 函数也是一个 goroutine

- `runtime.Goexit()` 结束当前的 goroutine，但不会产生 panic

- `runtime.Gosched()` 放弃当前调度，在队列中等待下次调度

## channel 通信和同步

[示例代码](./channel/channel.go)

- 不要通过共享内存来通信，而是通过通信来共享内存

- 通道用于 goroutine 通信，无缓冲通道用来同步

## WaitGroup 多个通道通信

[示例代码](./waitgroup/waitgroup.go)

```go

type WaitGroup struct {
    // contains filterd or unexported fields
}
// 添加等待信号
func (wg *WaitGroup) Add(delta int)
// 释放等待信号
func (wg *WaitGroup) Done()
// 等待
func (wg *WaitGroup) Wait()

```
