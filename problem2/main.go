package main
import "fmt"

func robotPath(x int, y int) int {
	ways:=make([][]int,x)

	for i := range ways {
        ways[i] = make([]int, y)
    }
	for i:=0;i<x;i++{
		ways[i][0]=1
	}

	for j:=0;j<y;j++{
		ways[0][j]=1
	}
	for i:=1;i<x;i++{
		for j:=1;j<y;j++{
			ways[i][j]=ways[i-1][j]+ways[i][j-1]
		}
	}
	return ways[x-1][y-1]
}

func main(){
	fmt.Println(robotPath(4,4))
}
