package main

// https://www.liwenzhou.com/posts/Go/interface/#c-1-3-2

// Interface 定义通过索引对元素排序的接口类型
type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

// reverse 结构体中嵌入了Interface接口
type reverse struct {
	Interface
}

/*  通过在结构体中嵌入一个接口类型，
从而让该结构体类型实现了该接口类型，
并且还可以改写该接口的方法。
*/

// Less 为reverse类型添加Less方法，重写原Interface接口类型的Less方法
func (r reverse) Less(i, j int) bool {
	return r.Interface.Less(j, i)
}

/* Interface类型原本的Less方法签名为 Less(i, j int) bool，
此处重写为r.Interface.Less(j, i)，即通过将索引参数交换位置实现反转。

在这个示例中还有一个需要注意的地方是reverse结构体本身是不可导出的（结构体类型名称首字母小写）
，sort.go中通过定义一个可导出的Reverse函数来让使用者创建reverse结构体实例。 */

func Reverse(data Interface) Interface {
	return &reverse{data}
}

/* 这样做的目的是保证得到的reverse结构体中的Interface属性一定不为nil，
否者r.Interface.Less(j, i)就会出现空指针panic。 */

func main() {
	r := &reverse{}
	r.Interface.Len()
	r.Len()
}
