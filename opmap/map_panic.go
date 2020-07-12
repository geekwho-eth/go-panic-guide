package opmap

import (
	"fmt"
	"sync"
)

// panic_add_key_to_nil_map
func addKey2NilMap() {
	var m map[string]int

	elem := 2
	fmt.Println("Add a key-element pair to a nil map...")
	m["two"] = elem // 这里会引发panic。
}

// panic_with_map_interface
func mapInterface() {
	var panicMap = map[interface{}]int{
		"1":      1,
		[]int{2}: 2, // 这里会引发panic。
		3:        3,
	}
	_ = panicMap
}

func concurrentMapPanic() {
	m := make(map[int]int)
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			m[i] = i // panic: concurrent map writes
		}(i)
	}
	wg.Wait()
}
