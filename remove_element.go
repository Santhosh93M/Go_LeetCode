package main

import "fmt"

func removeElement(nums []int, val int) int {
    begin := 0
    for i:=0;i<len(nums);i++{
        if nums[i]!=val{
            nums[begin] = nums[i]
            begin++
        }
    }
    return begin
    
}

func main(){
	nums := []int {0,1,2,2,3,0,4,2}
	fmt.Println(removeElement(nums,2))
}