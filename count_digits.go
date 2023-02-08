package main 

import "fmt"

func countDigits(num int) int {
    count := 0
    temp := num
    for temp>0{
        rem := temp%10
		if rem==0{
			continue
		}
        if num%rem==0{
            count++
        }
        temp /= 10
    }
    return count
}

func main(){
	fmt.Println(countDigits(121))
	fmt.Println(countDigits(13420))
	
}
