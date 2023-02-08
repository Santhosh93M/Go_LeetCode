package main

import "fmt"

func runningSum(nums []int) []int {
    res := nums[0]
    for i:=0;i<len(nums);i++{
        if i==0{
            continue
        }else{
            res += nums[i]
            nums[i] = res
        }
        
    }
    return nums
    
}

func main(){
	nums := []int {1,2,3,4}
	fmt.Println(runningSum(nums))

}