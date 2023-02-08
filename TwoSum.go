package main

import "fmt"

func twoSum(nums []int, target int) []int {
    for i,val1:=range nums{
        for j,val2:=range nums{
            if i==j{
                continue
            }
            if val1+val2==target{
                return []int {i,j}
            }
        }
    }
    return []int{}
    
}

func main(){
	nums := []int {2,7,11,9,4}
	fmt.Println(twoSum(nums,9))
}