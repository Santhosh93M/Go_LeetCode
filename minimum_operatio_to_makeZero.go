package main

import "fmt"


func minimumOperations(nums []int) int {
   res :=make(map[int]int)
   for _,val:=range nums{
       if val!=0{
           res[val]++
       }
   }
   return len(res)
}

func main(){
	nums := []int {1,5,0,3,5}
	nums1 := []int {0}
	fmt.Println(minimumOperations(nums))
}