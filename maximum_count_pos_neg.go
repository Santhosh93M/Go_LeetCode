package main

import "fmt"

func maximumCount(nums []int) int {
    pos := 0
    neg := 0
    for _,j:=range nums{
        if j<0{
            neg++
        }else if j>0{
            pos++
        }
    }
    return max(pos,neg)
}
func max(num1 int,num2 int)int{
    if num1>num2{
        return num1
    }
    return num2
}

func main(){
	nums := []int {-2,-1,-1,1,2,3}
	fmt.Println(maximumCount(nums))
}