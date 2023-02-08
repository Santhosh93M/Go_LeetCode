package main

import "fmt"

func isSameAfterReversals(num int) bool {
    if num==0{
        return true
    }
    rev1 := reverse(num)
    rev2 := reverse(rev1)
    if rev2==num{
        return true
    }
    return false
}
func reverse(i int)int {
    rev := 0
    for i>0{
        remainder := i % 10
        rev = rev*10 + remainder
        i /= 10
    }
    return rev
}

func main(){
	fmt.Println(isSameAfterReversals(121))
	fmt.Println(isSameAfterReversals(12100))

}