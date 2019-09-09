package main

import (
	"fmt"
)

func main() {

	//Take two slice of some different values, then copy one to another
	arr0 := []int{1, 2, 3, 4, 5}
	arr1 := []int{6, 7, 8, 9}

	copy1 := copy(arr0, arr1) //will return number of items copied in the array
	fmt.Println(copy1, arr0)

}
