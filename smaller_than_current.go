package main

import "fmt"

func smallerNumbersThanCurrent(nums []int) []int {
    res := []int {}
    for i:=0;i<len(nums);i++{
        count := 0
        for j:=0;j<len(nums);j++{
            if nums[i]>nums[j]{
                count += 1
            }
        }
        res = append(res, count)
    }
    return res
    
}


func main(){
	nums := []int {8,1,2,2,3}
	fmt.Println(smallerNumbersThanCurrent(nums))
	

}