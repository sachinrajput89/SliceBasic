package main

import (
	"bytes"
	"fmt"
)

func main() {

	//Take two byte slices

	slice0 := []byte{'a', 'b', 'c', 'd', 'e', 'f'}
	slice1 := []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G'}
	slice2 := []byte{'a', 'b', 'c', 'd', 'e', 'f'}

	res := bytes.Compare(slice0, slice1)
	res1 := bytes.Compare(slice1, slice2)
	res2 := bytes.Compare(slice2, slice0)

	fmt.Println(res)
	fmt.Println(res1)
	fmt.Println(res2)

	//some other format to check equality

	res3 := bytes.Equal(slice0, slice1)
	res4 := bytes.Equal(slice1, slice2)
	res5 := bytes.Equal(slice2, slice0)

	fmt.Println(res3, res4, res5)
}
