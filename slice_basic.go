package main

import (
	"fmt"
	"sort"
)

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

	//=======================================================

	//few more slices from array

	myarr := [...]string{"a", "b", "c", "d", "e", "f", "g", "h"}

	myslice0 := myarr[2:]
	myslice1 := myarr[:3]
	myslice2 := myarr[:]

	fmt.Println(myslice0)
	fmt.Println(myslice1)
	fmt.Println(myslice2)
	//===========================================================

	//using range in slice
	myrangeslice := []string{"i", "j", "k", "l", "m"}

	for index, ele := range myrangeslice {

		fmt.Println(index, ele)
	}

	//===============================================================

	//sorting a slice

	unsortarray := []string{"r", "l", "z", "u", "e", "g", "a"}
	unsortedint := []int{3, 5, 6, 1, 7, 8}

	fmt.Println(unsortarray)
	fmt.Println(unsortedint)

	sort.Strings(unsortarray)
	sort.Ints(unsortedint)

	fmt.Println(unsortarray)
	fmt.Println(unsortedint)

}
