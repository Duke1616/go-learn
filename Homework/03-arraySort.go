package main

import (
	"fmt"
)

func main() {
	array := [...]int{3, 6, 7, 40, 2, 5, 9, 20, 6}
	//fmt.Println
	fmt.Println(len(array) - 1)
	temp := 0
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-1-i; j++ {
			if array[j] > array[j+1] {
				temp = array[j]
				array[j] = array[j+1]
				array[j+1] = temp
			}
		}
		fmt.Println(array)
	}
}
