package main

import "fmt"

// simple way
func fibonacci(x int) int {
	if x == 0 {
		return 0
	} else if x <= 2 {
		return 1
	} else {
		return fibonacci(x-2) + fibonacci(x-1)
	}
}

// memoization
type Memoized func(int) int

var fibMem Memoized

func memoize(mf Memoized) Memoized {
	cache := make(map[int]int)
	return func(key int) int {
		if val, found := cache[key]; found {
			return val
		}
		temp := mf(key)
		cache[key] = temp
		return temp
	}
}

func FibMemoized(n int) int {
	fibMem = memoize(func(n int) int {
		if n == 0 || n == 1 {
			return n
		}
		return fibMem(n-2) + fibMem(n-1)
	})
	return fibMem(n)
}

func main() {
	fmt.Println(fibonacci(10))
	fmt.Println(FibMemoized(10))
}
