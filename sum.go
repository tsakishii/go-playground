package main

func Sum(arr []int) int {
	sum := 0
	for i := range arr {
		sum += arr[i]
	}
	return sum
}
