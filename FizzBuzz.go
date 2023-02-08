package main

import "fmt"


func fizzBuzz(n int) []string {
    res := []string{}
    for i:=1;i<n+1;i++{
        if i%3==0&&i%5==0{
            res = append(res, "FizzBuzz")
        }else if i%3==0&&i%5!=0{
            res = append(res, "Fizz")
        }else if i%3!=0&&i%5==0{
            res = append(res, "Buzz")
        }else{
            res = append(res, fmt.Sprintf("%d",i))
        }

    }
    return res
    
}

func main(){
	num := 5
	fmt.Println(fizzBuzz(num))

}