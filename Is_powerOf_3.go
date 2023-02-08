package main

import "fmt"

func isPowerOfThree(n int) bool {
    if n==1{
        return true
    }
    count := 3
    for{
        if count==n{
            return true
        }else if count>n{
            return false
        }
        count *= 3
    }
    return false
    
}

func main(){
	fmt.Println(isPowerOfThree(9))
	fmt.Println(isPowerOfThree(45))

}