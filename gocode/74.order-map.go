package main

import (
	"container/list"
	"fmt"
)

// 利用 Map 和 List 实现有序Map

type OrderMap interface {
	Exists(data Keyer) bool
	Push(data Keyer) bool
	Remove(data Keyer)
	Size() int
	Walk(cb func(data Keyer))
}

type Keyer interface {
	GetKey() string
}

// 用两种结构，保存同一份数据。
// 用 list 做顺序遍历，map 用来做查找，删除
type MapList struct {
	dataMap  map[string]*list.Element
	dataList *list.List
}

func NewMapList() *MapList {
	return &MapList{
		dataMap:  make(map[string]*list.Element),
		dataList: list.New(),
	}
}

func (mapList *MapList) Exists(data Keyer) bool {
	_, exists := mapList.dataMap[string(data.GetKey())]
	return exists
}

func (mapList *MapList) Push(data Keyer) bool {
	if mapList.Exists(data) {
		return false
	}
	elem := mapList.dataList.PushBack(data)
	mapList.dataMap[data.GetKey()] = elem
	return true
}

func (mapList *MapList) Remove(data Keyer) {
	if !mapList.Exists(data) {
		return
	}
	mapList.dataList.Remove(mapList.dataMap[data.GetKey()])
	delete(mapList.dataMap, data.GetKey())
}

func (mapList *MapList) Size() int {
	return mapList.dataList.Len()
}

func (mapList *MapList) Walk(cb func(data Keyer)) {
	for elem := mapList.dataList.Front(); elem != nil; elem = elem.Next() {
		cb(elem.Value.(Keyer))
	}
}

type Elements struct {
	value string
}

func (e Elements) GetKey() string {
	return e.value
}

func main() {
	fmt.Println("Starting test...")
	ml := NewMapList()
	// var a,b,c Keyer
	a := &Elements{"Alice"}
	b := &Elements{"Bob"}
	c := &Elements{"Conrad"}
	ml.Push(a)
	ml.Push(b)
	ml.Push(c)
	cb := func(data Keyer) {
		fmt.Println(ml.dataMap[data.GetKey()].Value.(*Elements).value)
	}
	fmt.Println("Pront elements in the order of pushing:")
	ml.Walk(cb)
	fmt.Printf("Size of MapList: %d \n", ml.Size())
	ml.Remove(b)
	fmt.Println("After removing b:")
	ml.Walk(cb)
	fmt.Printf("Size of MapList: %d \n", ml.Size())
}
