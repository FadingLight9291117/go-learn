package main

func main() {
	syncByChan()
}

// 无缓冲通道用于同步
func syncByChan() {
	c := make(chan struct{})
	go func() {
		sum := 0
		for i := 0; i < 1000; i++ {
			sum += i
		}
		println("end")
		c <- struct{}{}
	}()
	<-c
}
