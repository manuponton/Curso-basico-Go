package main

import "fmt"

func main() {
	var array [4]int
	array[0] = 1
	array[1] = 2
	fmt.Println(array, len(array), cap(array))

	slice := []int{0, 1, 2, 3, 5, 6, 7}
	fmt.Println(slice, len(slice), cap(slice))

	slice = append(slice, 8)
	fmt.Println(slice)
	newSlice := []int{9, 10, 11}
	slice = append(slice, newSlice...)
	fmt.Println(slice)
}
