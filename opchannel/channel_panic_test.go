package opchannel

import (
	"log"
	"testing"
)

func Test_closeClosedChannel(t *testing.T) {
	// 从 panic 恢复
	defer func() {
		if r := recover(); r != nil {
			// 成功从 panic 恢复
			log.Println("Test passed, panic was caught!")
		}
	}()
	// 调用函数触发 panic
	closeClosedChannel()
}

func Test_sendData2CloseChannel(t *testing.T) {
	// 从 panic 恢复
	defer func() {
		if r := recover(); r != nil {
			// 成功从 panic 恢复
			log.Println("Test passed, panic was caught!")
		}
	}()
	// 调用函数触发 panic
	sendData2CloseChannel()
}
