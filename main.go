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
}
