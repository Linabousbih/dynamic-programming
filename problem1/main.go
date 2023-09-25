package main
import "fmt"

func strairs(n int) int {
	// var ways int
	remaining:=map[int]int{0:1,1:1,2:2}
	if remaining[n]!=0{
		return remaining[n]
	} else{
		remaining[n]=strairs(n-1)+strairs(n-2)+strairs(n-3)
	}
	return remaining[n]
}

func main(){
	fmt.Println(strairs(5))
}