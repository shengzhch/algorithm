package main

import "fmt"

func main() {
	fmt.Println(isValid("(()))((()))"))
}

func isValid(s string) (int, bool) {
	var cmax, rmax int
	var wait int
	for _, r := range s {
		switch r {
		case '(':
			wait++
		case ')':
			if wait > 0 {
				wait--
				cmax = cmax + 2
				if cmax > rmax {
					rmax = cmax
				}
			}
			if wait == 0 {
				cmax = 0
			}
		}
	}
	return rmax, rmax == len(s)
}
