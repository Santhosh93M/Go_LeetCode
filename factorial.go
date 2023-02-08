package main

import (
	"fmt"
	"unicode"
)

func factorial(num int) int {
	fact := 1
	if num < 0 {
		return 0
	} else if num == 0 {
		return 1
	} else {
		for i := 1; i < num+1; i++ {
			fact *= i
		}
	}
	return fact
}

func main() {
	fmt.Println(factorial(3))
	for _, val := range "saTthosh" {
		fmt.Println(unicode.IsUpper(val))
	}

}
