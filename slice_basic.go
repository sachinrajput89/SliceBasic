package main

import "fmt"

func main() {

	//creating a basic slice operation

	slice0 := [...]int{1, 2, 3, 4, 5, 6, 7}

	fmt.Println(slice0)
	//slice at 1 and 3
	sliceon3 := slice0[1:3]
	fmt.Println(sliceon3)

	fmt.Println(len(sliceon3))
	fmt.Println(cap(sliceon3))

	//=================================================

	//Creating a shorthand slice

	slicea := []string{"a", "b", "c", "d", "e"}
	fmt.Println(slicea)

}
