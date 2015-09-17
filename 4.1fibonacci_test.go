package main

import (
	"fmt"
	"testing"
)

func TestFibo(t *testing.T) {
	var expected int = 13
	actual := fibonacci(7)
	if actual != expected {
		fmt.Println("The expected fibonacci is", actual)
		t.Error("Test failed")
	}
}

func fibonacci(x int) int {
	if x == 0 {
		return 0
	} else if x == 1 {
		return 1
	} else {
		return fibonacci(x-1) + fibonacci(x-2)
	}
}
func main() {
	x := int(7)
	fmt.Println(fibonacci(x))
}
