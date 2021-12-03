package main

import (
	"fmt"
)

func StateOfHealth(suggest *[3]string, i int, fatRate, a, b, c, d float64) {
	if fatRate <= a {
		fmt.Println("目前是：偏瘦")
		suggest[i] = "偏瘦"
	} else if fatRate > a && fatRate <= b {
		fmt.Println("目前是：标准")
		suggest[i] = "标准"
	} else if fatRate > b && fatRate <= c {
		fmt.Println("目前是：偏胖")
		suggest[i] = "偏胖"
	} else if fatRate > c && fatRate <= d {
		fmt.Println("目前是：肥胖")
		suggest[i] = "肥胖"
	} else {
		fmt.Println("目前是：非常肥胖")
		suggest[i] = "非常肥胖"
	}
}

func main() {
	weights := [3]float64{}
	names := [3]string{}
	bmis := [3]float64{}
	fatRates := [3]float64{}
	suggest := &[3]string{}

	for i := 0; i < 3; i++ {
		fmt.Printf("请输入第%d位同学的信息\n", i+1)
		var name string
		fmt.Print("名字：")
		fmt.Scanln(&name)
		names[i] = name

		var weight float64
		fmt.Print("体重（千克）: ")
		fmt.Scanln(&weight)
		weights[i] = weight

		var tall float64
		fmt.Print("身高（米): ")
		fmt.Scanln(&tall)

		var bmi float64 = weight / (tall * tall)
		bmis[i] = bmi

		var age int
		fmt.Print("年龄：")
		fmt.Scanln(&age)

		var sexWeight int
		var sex string
		fmt.Print("性别（男/女）：")
		fmt.Scanln(&sex)

		if sex == "男" {
			sexWeight = 1
		} else if sex == "女" {
			sexWeight = 0
		} else {
			// TODO
		}

		var fatRate float64 = (1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*float64(sexWeight)) / 100
		fatRates[i] = fatRate
		fmt.Println("体脂率是：", fatRate)

		if sex == "男" {
			// 编写男性的体脂率与体制状态表
			if age >= 18 && age <= 39 {
				StateOfHealth(suggest, i, fatRate, 0.1, 0.16, 0.21, 0.26)
			} else if age >= 40 && age <= 59 {
				StateOfHealth(suggest, i, fatRate, 0.11, 0.17, 0.22, 0.27)
			} else {
				StateOfHealth(suggest, i, fatRate, 0.13, 0.19, 0.24, 0.29)
			}
		} else {
			// 编写女性的体脂率与体制状态表
			if age >= 18 && age <= 39 {
				StateOfHealth(suggest, i, fatRate, 0.2, 0.27, 0.34, 0.39)
			} else if age >= 40 && age <= 59 {
				StateOfHealth(suggest, i, fatRate, 0.21, 0.28, 0.35, 0.40)
			} else {
				StateOfHealth(suggest, i, fatRate, 0.22, 0.29, 0.36, 0.41)
			}
		}
	}

	var sum float64
	for index, value := range fatRates {
		fmt.Printf("姓名：%s  BMI: %f  体脂率：%f 建议: %s \n", names[index], bmis[index], fatRates[index], suggest[index])
		sum += value
	}
	sum = sum / float64(len(fatRates))
	fmt.Printf("人数：%d   平均体脂: %f  \n", len(names), sum)
}
