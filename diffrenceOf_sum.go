package main

import "fmt"

func differenceOfSum(nums []int) int {
    sum := 0
    for _,val:= range nums{
        sum+=val
    }
    return sum-sumdigit(nums)
}
func sumdigit(num []int)int{
    sum:=0
    for _,val:=range num{
        for val>0{
            rem:=val%10
            sum+=rem
            val /= 10
        }
    }
    return sum
}

func main(){
	nums := []int {1,15,6,3}
	fmt.Println(differenceOfSum(nums))

}