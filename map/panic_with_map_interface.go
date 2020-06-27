package main

func main() {
	var panicMap = map[interface{}]int{
		"1":      1,
		[]int{2}: 2, // 这里会引发panic。
		3:        3,
	}
	_ = panicMap
}

