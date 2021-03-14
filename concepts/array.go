package main

import "fmt"

/**
 Slices: An array has fixed size. A slice, on the other hand, is flexible view into the elements of an array. It's more common than array, in practice.
 */

func makeSlices() {
	// static array declaration
	arr := [6]int{1,2,3,4,5,6}
    // making the slices
    // slices are reference to an array i.e if you change the resulted slice it will change the original array
	firstFourElements := arr[:4]
	fmt.Println("=============first four slices=========")
	fmt.Println(firstFourElements)

	// slices literals: like array without the length
	// it creates the array and builds the slice that reference it
	/**
	  length: length denotes the length of the array
	  capacity: capacity denotes the total number of element present in the array
	 */

	slicesLiterals := []bool {true , false}
	fmt.Println(slicesLiterals)
}

func createDynamicSizeOfArray(){
	// make: built-in function used to crate the dynamic size array
	a := make([]int, 5)
	fmt.Println(a)
}

func appendElementInSlice(){
	var s []int
	// append: builtin function used to append the element in the slices
	s = append(s,1)
	s = append(s, 1,2,3,4)
	fmt.Println(s)
}

/**
  The range form of the for loop iterates over a slice or map.
 */
func rangeF() {
	var pow = []int {1, 2, 4, 8, 16}

	for i , v := range pow {
		fmt.Printf("2**%d = %d\n",i, v)
	}
}

/**
 two dimensional array declaration
 */

func twoDArray(n, m int) {
	fmt.Println("===========two dimensional array==========")
	var array = make([][]int, n)
	for i := range array {
		array[i] = make([]int, m)
	}

	for i := range array {
		for j := range array[i] {
			fmt.Print(array[i][j], " ")
		}
		fmt.Print("\n")
	}
}

func main(){
	makeSlices()
	createDynamicSizeOfArray()
	appendElementInSlice()
	rangeF()
	twoDArray(2,4)

}