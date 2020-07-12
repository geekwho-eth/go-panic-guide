package typeassertion

import "fmt"

func typeAssertionPanic() {
	var i interface{} = "hello"
	n := i.(int) // panic: interface conversion: string is not int
	fmt.Println(n)
}

func typeAssertionWork(i interface{}) int {
	if v, ok := i.(int); ok {
		return v
	} else {
		return 0
	}
}
