package main

import (
	"runtime"
	"time"
)

func main() {
	println("GOMAXPROCS=", runtime.GOMAXPROCS(0)) // 默认为核心数
	runtime.GOMAXPROCS(1)
	for i := 0; i < 1000000; i++ {
		go Sum()
	}
	println("Num of Goroutine is ", runtime.NumGoroutine()) // runtime.NumGroutine() 返回当前程序的goroutine数目
	println("GOMAXPROCES=", runtime.GOMAXPROCS(0))
	time.Sleep((5 * time.Second))
}

func Sum() {
	sum := 0
	for i := 0; i < 10000; i++ {
		sum += i
	}
	time.Sleep(1 * time.Second)
}
