package main

import "fmt"

func commonFactors(a int, b int) int {
    min := 0
    count := 0
    if a>b{
        min = b
    }else{
        min = a
    }
    for i:=1;i<min+1;i++{
        if a%i==0 && b%i==0{
            count++
        }
    }
    return count
    
}

func main(){
	fmt.Println(commonFactors(12,6))
	fmt.Println(commonFactors(25,30))
	
}