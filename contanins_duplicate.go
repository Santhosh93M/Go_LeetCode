package main

import "fmt"

func containsDuplicate(nums []int) bool {
    for j,i:=range nums{
        for k:=j+1;k<len(nums);k++{
            if i==nums[k]{
                return true
            }
        }
    }
    return false
}

func main(){
	nums := []int {1,2,3,1}
	nums1 := []int {1,2,3,4}
	fmt.Println(containsDuplicate(nums))
	fmt.Println(containsDuplicate(nums1))

}