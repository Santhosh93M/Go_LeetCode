package main

import "fmt"

func isPowerOfTwo(n int) bool {
    if n==1{
        return true
    }
    count := 2
    for{
        if count==n{
            return true
        }else if count>n{
            return false
        }
        count *= 2
    }
    return false
    
}

func main(){
	fmt.Println(isPowerOfTwo(8))
	fmt.Println(isPowerOfTwo(10))

}