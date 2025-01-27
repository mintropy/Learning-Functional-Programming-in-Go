package main

import "fmt"

func MyChannel(ch chan int, counter int) {
	n1, n2 := 0, 1
	for i := 0; i < counter; i++ {
		ch <- n1
		n1, n2 = n2, n1+n2
	}
	close(ch)
}

func FibChanneled(n int) int {
	n += 2
	ch := make(chan int)
	go MyChannel(ch, n)
	i := 0
	var result int
	for num := range ch {
		result = num
		i++
	}
	return result
}

func main() {
	fmt.Println(FibChanneled(10))
}
