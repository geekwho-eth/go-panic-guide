package oparray

import "fmt"

func sliceOutOfRangePanic() {
	//arr := [3]int{1, 2, 3}
	arr := make([]int, 3)
	fmt.Println(arr[3]) // panic: index out of range
}
