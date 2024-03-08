package moretypes

import "fmt"

// Interface 作为万能数据类型

func MoreTypes() []interface{} {
	var a1 interface{} = 1
	var a2 interface{} = "abc"
	list := make([]interface{}, 0)
	list = append(list, a1)
	list = append(list, a2)
	fmt.Println(list)
	return list
}

func hello() {
	fmt.Println("")
}
