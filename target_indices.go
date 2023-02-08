package main

import "fmt"

func targetIndices(nums []int, target int) []int {
    res := []int {}
    sort.Ints(nums)
    for i,j:=range nums{
        if j==target{
            res = append(res, i)
        }
    }
    return res
    
}

func main(){
	nums := []int {1,2,5,2,3}
	fmt.Println(targetIndices(nums))

}