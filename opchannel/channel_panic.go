package opchannel

import "fmt"

func closeClosedChannel() {
	ch1 := make(chan int, 3)
	ch1 <- 2
	ch1 <- 1
	ch1 <- 3
	elem1 := <-ch1
	fmt.Printf("The first element received from channel ch1: %v\n",
		elem1)
	close(ch1)
	close(ch1) // go panic
}

func sendData2CloseChannel() {
	ch1 := make(chan int, 3)
	ch1 <- 2
	ch1 <- 1
	ch1 <- 3
	elem1 := <-ch1
	fmt.Printf("The first element received from channel ch1: %v\n",
		elem1)
	close(ch1)
	// 发送数据到已关闭的channel，会产生panic
	ch1 <- 4
}

func createNoBufferedChannelPanic() {
	ch := make(chan int) // 创建一个无缓冲的通道
	ch <- 1              // 主 goroutine 在这里阻塞等待接收者，但没有接收者 —— panic
}

func createBufferedChannelPanic() {
	ch := make(chan int, 1)
	ch <- 1
}

func fixedChannelPanic() {
	ch := make(chan int)

	go func() {
		fmt.Println("Received:", <-ch)
	}()

	ch <- 1 // 正常发送，不再 panic
}

func fixedBufferedChannel() {
	ch := make(chan int, 1)
	ch <- 1 // 不会阻塞，因为 channel 有一个缓冲位
	fmt.Println("Sent to buffered channel")
}
