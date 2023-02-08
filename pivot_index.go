package main

import "fmt"

func pivotIndex(nums []int) int {
    for ind,_:=range nums{
        res := getleftrightsum(nums,ind)
        fmt.Println(res)
        if res[0]==res[1]{
            return ind
        }
    }
    
    return -1
}
func getleftrightsum(lst []int,key int)[]int{
    sum1:=0
    sum2:=0

    for i:=0;i<key;i++{
        sum1+=lst[i]
    }

    for j:=key+1;j<len(lst);j++{
        sum2+=lst[j]
    }

    return []int{sum1,sum2}

}

func main(){
	nums := []int {1,7,3,6,5,6}
	nums1 := []int {1,2,3}
	fmt.Println(pivotIndex(nums))
	fmt.Println(pivotIndex(nums1))

}