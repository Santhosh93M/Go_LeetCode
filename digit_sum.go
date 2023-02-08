package main

import "fmt"

func addDigits(num int) int {

    sum := 0

	for num > 0 {
		sum += num % 10
		num /= 10
	}

	return sum
    
}

func main(){
	fmt.Println(addDigits(12235654))
	
}