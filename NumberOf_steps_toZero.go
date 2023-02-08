package main

import "fmt"

func numberOfSteps(num int) int {
    if num==0{
        return 0
    }
    count := 0
    for num > 0{
        if num%2==0{
            num /= 2
            count+=1
        }else{
            num -= 1
            count += 1
        }
    }
    return count
    
}

func main(){
	fmt.Println(numberOfSteps(14))
}