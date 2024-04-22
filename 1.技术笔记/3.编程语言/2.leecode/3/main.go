package main

import (
	"fmt"
)

/*
给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串的长度。

示例 1:
输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

示例 2:
输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。

示例 3:
输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串

提示：
0 <= s.length <= 5 * 104
s 由英文字母、数字、符号和空格组成
*/

// 滑动窗口解法
func lengthOfLongestSubstring(s string) int {

	//使用 map作为不重复元素判断
	m := make(map[byte]int)
	//字符串长度
	n := len(s)
	//右指针和子串最大长度
	rk, ack := -1, 0
	for i := 0; i < n; i++ {
		//如果不是第一个元素，那么每次滑动窗口时都要弹出最左侧的元素
		if i != 0 {
			delete(m, s[i-1])
		}
		//如果没超出字符串长度或者map[s[rk+1]]==0，就向右移动右指针
		for rk+1 < n && m[s[rk+1]] == 0 {
			//给 map[s[rk+1]] 赋值 + 1，代表着该字符已经出现过一次
			m[s[rk+1]]++
			//指针右移
			rk++
		}
		//找出最长子串长度
		ack = max(ack, rk-i+1)
	}
	return ack
}

// 取两个数中的大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {

	str := "hello"
	fmt.Println(str[1])
}
