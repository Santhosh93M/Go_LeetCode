package main

import "fmt"

func kidsWithCandies(candies []int, extraCandies int) []bool {
    res := []bool{}
    max := 0
    for _,i:=range candies{
        if max < i{
            max = i
        }
    }
    for _,j:= range candies{
        if j+extraCandies >= max{
            res = append(res, true)
        }else{
            res = append(res, false)
        }
    }
    return res
    
}

func main(){
	nums := []int {2,3,5,1,3}
	fmt.Println(kidsWithCandies(nums,3))

}