package main

import "fmt"

func countNegatives(grid [][]int) int {
    neg := 0
    for _,i:=range grid{
        for _,j:=range i{
            if j<0{
                neg++
            }
        }
    }
    return neg
}

func main(){
	nums := [][]int {{4,3,2,-1},{3,2,1,-1},{1,1,-1,-2},{-1,-1,-2,-3}}
	fmt.Println(countNegatives(nums))
}