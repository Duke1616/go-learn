package main

import (
	"fmt"
)

func main() {
	array := [...]string{"栾凯朝", "LISI", "ZhangSan"}
	for index, value := range array {
		chineseBytes := len([]rune(value))
		abytes := []rune(value)
		intv := len(value)
		// 判断如果是中文字符串不做处理
		if intv == chineseBytes {
			for i, v := range abytes {
				// 判断字符串是大写 OR 小写
				if v > 96 && v < 123 {
					abytes[i] = v - 32
					value = string(abytes)
					array[index] = value
				} else if v > 64 && v < 123 {
					abytes[i] = v + 32
					value = string(abytes)
					array[index] = value
				} else {
					fmt.Println("不做任何处理")
				}
			}
		}
	}
	fmt.Println(array)
}
