package main

import "fmt"

func main() {
	var m map[string]int

	elem := 2
	fmt.Println("Add a key-element pair to a nil map...")
	m["two"] = elem // 这里会引发panic。
}
