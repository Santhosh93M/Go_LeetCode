package main

import "fmt"

func searchInsert(nums []int, target int) int {

    for i,val:=range nums{
        if val == target || val > target{
            return i
        }
    }
    return len(nums)
}

func main(){
	nums := []int {1,3,5,6}
	fmt.Println(searchInsert(nums,5)
	fmt.Println(searchInsert(nums,2)
	fmt.Println(searchInsert(nums,7)

}