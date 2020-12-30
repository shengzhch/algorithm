package basic

import (
	"fmt"
	"reflect"
)

var CfInt = func(key1, key2 interface{}) int {
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

var HashInt = func(key interface{}) int {
	return key.(int)
}

var MatchInt = func(key1, key2 interface{}) bool {
	return reflect.DeepEqual(key1, key2)
}

func ToInterfaceSlice(in interface{}) []interface{} {
	if reflect.TypeOf(in).Kind() != reflect.Slice {
		return []interface{}{in}
	}
	val := reflect.ValueOf(in)
	l := val.Len()
	rel := make([]interface{}, 0, l)
	for i := 0; i < l; i++ {
		rel = append(rel, val.Index(i).Interface())
	}
	return rel
}
