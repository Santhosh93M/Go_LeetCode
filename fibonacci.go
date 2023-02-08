package main

import "fmt"

func fibonacci(n int) {
	a := 0
	b := 1
	c := b
	fmt.Printf("Fibonacci Series : %d %d ", a, b)
	for true {
		c = b
		b = a + b
		if b >= n {
			fmt.Println()
			break
		}
		a = c
		fmt.Printf("%d ", b)
	}
}

func main() {
	fibonacci(10)
}
