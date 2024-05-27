package main

import "fmt"

func main() {
	array := [3]int{1, 2, 3}
	var slice []int

	//array = append(array, 4, 5, 6)
	slice = append(slice, 7, 8, 9)
	fmt.Println(array, slice)

	slice2 := make([]int, 2)
	copy(slice2, slice)
	fmt.Println(slice2)
}
