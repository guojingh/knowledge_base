package main

import "fmt"

// 输出最长子字符串的内容
func main() {

	fmt.Println(getMaxSonString("oooooooHelloaaaccccc"))
}

func getMaxSonString(s string) string {
	bytes := []byte(s)
	maxLen := 1
	tmpLen := 1
	var tmp byte
	var index = 0
	for i := 0; i < len(bytes); i++ {
		if i == 0 {
			tmp = bytes[i]
			continue
		}
		if tmp == bytes[i] {
			tmpLen++
			if tmpLen > maxLen {
				maxLen = tmpLen
				index = i - maxLen + 1
			}
		} else {
			tmpLen = 1
		}
		tmp = bytes[i]
	}

	return string(bytes[index : index+maxLen])
	//return maxLen
}
