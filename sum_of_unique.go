package main

import "fmt"

func sumOfUnique(nums []int) int {
    if len(nums)==1{
        return nums[0]
    }
    hash := make(map[int]int)
    sum := 0
    for _,val:=range nums{
        hash[val]++
    }
    for key,val := range hash{
        if val==1{
            sum+=key
        }
    }
    return sum
}

func main(){
	nums := []int {1,2,2,3}
	fmt.Println(sumOfUnique(nums))

}