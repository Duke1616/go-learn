package main

import "fmt"

func main() {
	var x11 float64 = 20
	var x21 float64 = 40
	var y11 float64 = 60
	var y21 float64 = 70

	var x12 float64 = 30
	var x22 float64 = 40
	var y12 float64 = 50
	var y22 float64 = 20

	var k1 float64 = (y21 - y11) / (x21 - x11)
	var k2 float64 = (y22 - y12) / (x22 - x12)

	if k1 == k2 {
		fmt.Println("两条平行")
	} else {
		fmt.Println("两条不平行")
	}
}
