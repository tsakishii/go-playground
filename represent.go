package main

import "fmt"

func Represent(n [][]int) {
	for i := 0; i <= len(n)-1; i++ {
		for j := 0; j <= len(n)-1; j++ {
			fmt.Printf("%d ", n[i][j])
		}
		fmt.Println()
	}
}
