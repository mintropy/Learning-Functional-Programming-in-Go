package main

import "fmt"

// named function
func named(s string) {
	fmt.Println(s)
}

// anonymous function
func anonymous() func(string) {
	return func(s string) {
		fmt.Println(s)
	}
}

// closure
func closure() func() {
	s := "closure"
	_closure := func() {
		fmt.Println(s)
	}
	return _closure
}

func main() {
	named("named function")
	anonymous()("anonymous function")
	closure()()
}
