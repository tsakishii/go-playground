package main

import (
	"fmt"
	"encoding/json"
)

func main() {
	// create a multi dim slice
	some := [][]int{{1,2,3}, {4,5,6}, {7, 8, 9}}

	// put the slice in a chanel
	ch := make(chan [][]int, 1)
	go func() {
		ch <- some
		close(ch)
	}()

	mat:= <- ch
	Represent(mat)


	// funny rugby
	foo := Foo{"Joe Marler", "Harlequins Tigers"}
	but, _ := json.Marshal(foo)
	fmt.Println(string(but))
	json.Unmarshal(but, &foo)

	arr := [...]int{1, 2, 3, 4}
	sum := Sumar(&arr)

	fmt.Printf("sum of arr: %d added from array {1,2,3,4}\n", sum)

	arr1 := []int{1, 2, 3, 4}
	sumarr1 := Sum(arr1)
	fmt.Printf("sum of slice: %d\n", sumarr1)

	// Test factorials
	fmt.Println(Factorial(4.0))
	fmt.Println(FactorialIter(4.0))
}
