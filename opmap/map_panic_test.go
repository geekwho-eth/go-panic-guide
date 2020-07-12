package opmap

import (
	"log"
	"testing"
)

func Test_addKey2NilMap(t *testing.T) {
	// 从 panic 恢复
	defer func() {
		if r := recover(); r != nil {
			// 成功从 panic 恢复
			log.Println("Test passed, panic was caught!")
		}
	}()
	// 调用函数触发 panic
	addKey2NilMap()
}

func Test_mapInterface(t *testing.T) {
	// 从 panic 恢复
	defer func() {
		if r := recover(); r != nil {
			// 成功从 panic 恢复
			log.Println("Test passed, panic was caught!")
		}
	}()
	// 调用函数触发 panic
	mapInterface()
}

func Test_concurrentMapPanic(t *testing.T) {
	// 当多个 goroutine 并发读写原始 map 时，会触发运行时检测，导致 fatal error: concurrent map writes，程序直接崩溃，recover 无法挽救。
	// 调用函数触发 panic
	concurrentMapPanic()
}
