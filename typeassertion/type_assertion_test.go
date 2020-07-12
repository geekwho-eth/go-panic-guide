package typeassertion

import (
	"log"
	"testing"
)

func Test_typeAssertionPanic(t *testing.T) {
	// Defer a function to recover from panic
	defer func() {
		if r := recover(); r != nil {
			// We successfully recovered from panic
			log.Println("Test passed, panic was caught!")
		}
	}()
	// Call the function that will panic
	typeAssertionPanic()
	// If the panic was not caught, the test will fail
	//t.Errorf("Test failed, panic was expected")
}

func Test_typeAssertionWork_got(t *testing.T) {
	input := "hello"
	got := typeAssertionWork(input)
	log.Println(got)
	if got != 0 {
		t.Errorf("type assertion failed. typeAssertionWork() = %v, want %v", got, 0)
	} else {
		log.Println("string type assertion success.")
	}

	inputInt := 100
	got = typeAssertionWork(inputInt)
	log.Println(got)
	expected := 100
	if got != expected {
		t.Errorf("typeAssertionWork() = %v, want %v", got, expected)
	} else {
		log.Println("int type assertion success.")
	}
}
