package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(lengthOfLongestSubstring2("pwwkew"))
}

func lengthOfLongestSubstring(s string) int {
	var rel int
	var count = make(map[rune]struct{})
	var maxK int
	var indexK int
	var j int //start
	for i := 0; i < len(s); i++ {
		if i > 0 {
			maxK = maxK - 1
			delete(count, rune(s[i-1]))
		}
		for j = indexK; j < len(s); j++ {
			if _, ok := count[rune(s[j])]; !ok {
				count[rune(s[j])] = struct{}{}
				maxK++
			} else {
				indexK = j
				break
			}
		}
		if maxK > rel {
			rel = maxK
		}
	}
	return rel
}

func lengthOfLongestSubstring2(s string) int {
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		index := strings.Index(s[start:i], string(s[i]))
		if index == -1 {
			if i+1 > end {
				end = i + 1
			}
		} else { //窗口同时向右滑动
			start += index + 1
			end += index + 1
		}
	}
	return end - start
}

func lengthOfLongestSubstring3(s string) int {
	r, l := 0, 0
	var ret int
	for i := range s {
		index := strings.Index(s[l:i], string(s[i]))
		if index == -1 {
			r++
		} else {
			r = i + 1
			l += index + 1
		}
		ret = max(len(s[l:r]), ret)
	}
	return ret
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
