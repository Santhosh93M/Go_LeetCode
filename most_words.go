packge main

import (
	"fmt"
	"strings"
)

func mostWordsFound(sentences []string) int {
    max := 0
    for _,sen := range sentences{
        if len(strings.Split(sen," ")) > max{
            max = len(strings.Split(sen," "))
        }
    }
    return max
}

func main(){
	sentences1 := []string {"alice and bob love leetcode", "i think so too", "this is great thanks very much"}
	sentences2 := []string {"please wait", "continue to fight", "continue to win"}
	
	fmt.Println(mostWordsFound(sentences1))

}