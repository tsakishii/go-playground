package main

import "fmt"

func represent(n [][]int) {
	for i := 0; i <= len(n)-1; i++ {
		for j := 0; j <= len(n)-1; j++ {
			fmt.Printf("%d ", n[i][j])
		}
		fmt.Println()
	}
}

func main() {
	// create a multi dim slice
	some := [][]int{{1,2,3}, {4,5,6}, {7, 8, 9}}
	represent(some)
}
