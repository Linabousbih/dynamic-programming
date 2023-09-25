package main

import "fmt"

func fibonacci(num int) int{
	fibonacciNumbers:=map[int]int{1:0,2:1}

	if val, ok := fibonacciNumbers[num]; ok {
        return val
    }else {
		fibonacciNumbers[num]=fibonacci(num-1)+fibonacci(num-2)
	}
	return fibonacciNumbers[num]
}

func main(){
	fmt.Println(fibonacci(9))
}