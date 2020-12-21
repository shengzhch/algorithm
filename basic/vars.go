package basic

import "fmt"

var CfInt= func(key1, key2 interface{}) int {
	if key1.(int) > key2.(int) {
		return 1
	} else if key1.(int) < key2.(int) {
		return -1
	} else {
		return 0
	}
}

var (
	PrintNode = func(e *Node, args ...interface{}) bool {
		if e == nil {
			return true
		}
		fmt.Println(e.Data())
		return false
	}
)
