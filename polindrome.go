package main

import "fmt"

func polindrome(num int) bool {
	temp := num
	res := 0
	for num > 0 {
		rem := num % 10
		res = (res * 10) + rem
		num /= 10
	}
	if temp == res {
		return true
	}
	return false
}

func main() {
	fmt.Println(polindrome(121))

}
