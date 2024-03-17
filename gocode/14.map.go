package main

/*
映射 (`map`)
映射将键映射到值。
映射的零值为 nil 。nil 映射既没有键，也不能添加键。
make 函数会返回给定类型的映射，并将其初始化备用。
*/
//也就是<K,V>键值对的形式

type Ver struct {
	Lat, Long float64
}

var m map[string]Ver

var m1 = map[string]Ver{ // 在赋值过程中，K的值不能为空
	"Bell Labs": {
		40.68433, -74.39967,
	},
	"Google": {
		37.42202, -122.08408,
	},
}

var m2 = map[string]Ver{ // 若顶级类型只是一个类型名，你可以在文法的元素中省略它。
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
}
