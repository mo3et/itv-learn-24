package main

import "fmt"

// 空结构体
var Exists = struct{}{}

// Set is the main interface
type Set struct {
	// struct 为结构体类型的变量
	m map[interface{}]struct{}
}

// 初始化
//
// Set类型数据的初始化操作，在声明的同时可以选择传入或者不传入。
// 声明Map切片的时候，Key是Interface{}类型的数据，Value用空struct实现。

func New(items ...interface{}) *Set {
	// 实例化Set
	s := &Set{}
	// 声明Set的数据结构 并实例化
	s.m = make(map[interface{}]struct{})
	s.Add(items...)
	return s
}

// 添加
//
// 使用变长参数来满足不定个数的元素传入
// 且Key值为唯一，不必排重，Value则用空结构体占位
func (s *Set) Add(items ...interface{}) error {
	for _, item := range items {
		s.m[item] = Exists
	}
	return nil
}

// 包含
// Contains操作其实就是查询操作，看看是否有对应的Item存在 利用map特性实现
// 因为不需要Value 值，可以通过ok来获取
func (s *Set) Contains(item interface{}) bool {
	_, ok := s.m[item]
	return ok
}

// Len 只需要调用len()获取Map长度即可：
func (s *Set) Size() int {
	return len(s.m)
}

// Clear 操作可以通过初始化覆盖Set来实现
func (s *Set) Clear() {
	s.m = make(map[interface{}]struct{})
}

// Equal
// 判断两个Set是否相等，通过遍历实现
// 将A的每个元素，在 B 中是否存在，若有一个不存在，则返回false退出
func (s *Set) Equal(other *Set) bool {
	// 如果Size不等 则退出
	if s.Size() != other.Size() {
		return false
	}

	// range 遍历 Contains查询每个值是否相同
	for key := range s.m {
		// 只要一个不存在就返回false
		if !other.Contains(key) {
			return false
		}
	}
	return true
}

// 子集 A为B的子集(包含在B里面)
func (s *Set) IsSubset(other *Set) bool {
	// s的size长于other(子集A大于父集), false
	if s.Size() > other.Size() {
		return false
	}

	// range 遍历每个key是否相同
	for key := range s.m {
		// 只要一个不存在就返回false
		if !other.Contains(key) {
			return false
		}
	}
	return true
}

func main() {
	// 创建两个集合
	set1 := New("apple", "banana", "cherry")
	set2 := New("banana", "cherry", "date")

	// 添加元素
	set1.Add("date")
	set2.Add("apple")

	// 检查元素是否存在
	fmt.Println("Does set1 contain 'date'?", set1.Contains("date"))
	fmt.Println("Does set2 contain 'apple'?", set2.Contains("apple"))

	// 检查集合的等价性
	fmt.Println("Are set1 and set2 equal?", set1.Equal(set2))

	// 检查集合的子集关系
	fmt.Println("Is set1 a subset of set2?", set1.IsSubset(set2))
}
