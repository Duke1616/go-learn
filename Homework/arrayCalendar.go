package main

import (
	"fmt"
)

/*
1 3 5 7 8 10 12
4 6 9 11
2
*/

func main() {
	//var month int
	var yeartime int
	fmt.Print("请输入年份：")
	fmt.Scanln(&yeartime)

	yearCalendar := [13][31]int{}
	for index1, value1 := range yearCalendar {
		num := 0
		for index2, _ := range value1 {
			num += 1
			yearCalendar[index1][index2] = num
		}
	}

	for index, _ := range yearCalendar {
		//fmt.Println(index+1, value)
		if index == 1 || index == 3 || index == 5 || index == 7 || index == 8 || index == 10 || index == 12 {
			fmt.Println(index, yearCalendar[index][0:])
		} else if index == 4 || index == 6 || index == 9 || index == 11 {
			fmt.Println(index, yearCalendar[index][0:30])
		} else if index == 2 {
			if yeartime%4 == 0 && yeartime%100 != 0 || yeartime%400 == 0 {
				fmt.Println(index, yearCalendar[index][0:29])
			} else {
				fmt.Println(index, yearCalendar[index][0:28])
			}
		}
	}
}
