package opmap

import "fmt"

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
