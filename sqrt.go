package main

import "fmt"

func mySqrt(x int) int {
    res := 0
    for i:=0;i<x;i++{
        if i*i==x{
            res = i
            break
        }else if i*i>x{
            res = i-1
            break
        }
    }
    return res
}

func main(){
	fmt.Println(mySqrt(8))
}