package main

/*
依照提示输入：年份 月份
程序自动输出指定月份的日历表
*/

import (
	"fmt"
)

//var weekday = [7]string{"周一", "周二", "周三", "周四", "周五", "周六", "周日"}
var weekday = [7]string{"Tu", "We", "Th", "Fr", "Sa", "Su", "Mo"}

func main() {
	var yeartime uint16
	fmt.Print("年份：")
	fmt.Scanln(&yeartime)

	var month uint16
	fmt.Print("月份：")
	fmt.Scanln(&month)

	Calendar(yeartime, month)
}

func Calendar(yeartime, month uint16) {
	var day uint16 = 1

	//fmt.Printf("%d年%d月%d日是:%d\n", yeartime, month, day, ZellerFunction2Week(yeartime, month, day))
	yearCalendar := &[13][31]int{}
	for index1, value1 := range yearCalendar {
		num := 0
		for index2, _ := range value1 {
			num += 1
			yearCalendar[index1][index2] = num
		}
	}

	if month == 1 || month == 3 || month == 5 || month == 7 || month == 8 || month == 10 || month == 12 {
		OutputYearMonth(yearCalendar, 31, yeartime, month, day)
	} else if month == 4 || month == 6 || month == 9 || month == 11 {
		OutputYearMonth(yearCalendar, 30, yeartime, month, day)
	} else if month == 2 {
		if yeartime%4 == 0 && yeartime%100 != 0 || yeartime%400 == 0 {
			OutputYearMonth(yearCalendar, 29, yeartime, month, day)
		} else {
			OutputYearMonth(yearCalendar, 28, yeartime, month, day)
		}
	}
}

func OutputYearMonth(yearCalendar *[13][31]int, days int, yeartime, month, day uint16) {
	var Org int
	if ZellerFunction2Week(yeartime, month, day) == 0 {
		Org = 6
	} else {
		Org = ZellerFunction2Week(yeartime, month, day) - 1
	}
	origin := 6 - Org
	fmt.Printf("        年份：%5d  月份: %3d \n", yeartime, month)
	for i, _ := range weekday {
		fmt.Printf("%5s", weekday[i])
	}
	fmt.Println()
	if origin == 0 {
		fmt.Printf("\t\t\t      ")
	} else if origin == 1 {
		fmt.Printf("\t\t\t ")
	} else if origin == 2 {
		fmt.Printf("\t\t    ")
	} else if origin == 3 {
		fmt.Printf("\t       ")
	} else if origin == 4 {
		fmt.Printf("\t  ")
	} else if origin == 5 {
		fmt.Printf("     ")
	}
	for i := 0; i < days; i++ {
		fmt.Printf("%5d", yearCalendar[month][i])
		if i == origin || i == origin+7 || i == origin+14 || i == origin+21 || i == origin+28 {
			//fmt.Println(origin)
			fmt.Println()
		}
	}
}

func ZellerFunction2Week(year, month, day uint16) int {
	var y, m, c uint16
	if month >= 3 {
		m = month
		y = year % 100
		c = year / 100
	} else {
		m = month + 12
		y = (year - 1) % 100
		c = (year - 1) / 100
	}

	week := y + (y / 4) + (c / 4) - 2*c + ((26 * (m + 1)) / 10) + day - 1
	if week < 0 {
		week = 7 - (-week)%7
	} else {
		week = week % 7
	}
	which_week := int(week)
	return which_week
}
