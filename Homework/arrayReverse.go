package main

import "fmt"

func main() {
	array := [...]int{1, 20, 30, 40, 50, 60, 7}
	arrayReverse := [...]int{1, 20, 30, 40, 50, 60, 7}
	for i, _ := range array {
		arrayReverse[i] = array[len(array)-i-1]
	}
	fmt.Println(array)
	fmt.Println(arrayReverse)
}
