package main

import "fmt"

func truncateSentence(s string, k int) string {
    lst := strings.Split(s," ")
    res := make([]string,k)
    for i:=0;i<k;i++{
        res[i] = lst[i]
    }   
    return strings.Join(res," ")
}

func main(){
	s := "Hello how are you Contestant"
	var k = 4
	fmt.Println(truncateSentence(s,k))
	
	
}