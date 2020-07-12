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
