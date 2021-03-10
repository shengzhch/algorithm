package main

import "fmt"

func main() {
	fmt.Println(reverseVowels("ai"))
}

// 和对称相关的，考虑 2 种对称遍历方式
// ... --l, mid, r++ ...	// 考虑各种边界问题
//     l++, ..., --r		// l <= r 相遇即边界
func reverseVowels(s string) string {
	if s == "" {
		return s
	}
	runes := []rune(s)
	n := len(runes)
	l, r := 0, n-1

	for l <= r {
		for l < r && !isVowel(runes[l]) {
			l++
		}
		for r > l && !isVowel(runes[r]) {
			r--
		}
		if isVowel(runes[l]) && isVowel(runes[r]) {
			runes[l], runes[r] = runes[r], runes[l]
		}
		l++
		r--
	}
	return string(runes)
}

func isVowel(r rune) bool {
	switch r {
	case 'a', 'A', 'e', 'E', 'i', 'I', 'o', 'O', 'u', 'U':
		return true
	default:
		return false
	}
}
