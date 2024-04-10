package skiplist

type Skiplist struct {
	head *node
	// level int
}

type node struct {
	nexts    []*node
	key, val int
}

// Read process
// 根据 key读取val, 第二个 bool flag 反映 key 在 skiplist 是否存在

func (s *Skiplist) Get(key int) (int, bool) {
	// 根据key 尝试检索对应的 node, 如果node 存在，则返回对应的 val
}
