package main

import "fmt"

func armstrong(num int) bool {
	temp := num
	res := 0
	for num > 0 {
		rem := num % 10
		res += (rem * rem * rem)
		num /= 10
	}
	if temp == res {
		return true
	}
	return false

}

func main() {
	fmt.Println(armstrong(153))
}
