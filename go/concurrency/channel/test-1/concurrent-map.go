package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
*	一道考题(小徐先生 Channel 考题)
*   要求实现一个map:
*	1. 面向高并发
*	2. 只存在插入和查询操作 O(1)
*	3. 查询时，若key 存在，直接返回val;若key不存在，阻塞直到key val 对被放入后
*	获取 val 返回；等待指定时长仍未放入，返回超时错误
*	4. 不能有死锁或panic
 */

type MyChan struct {
	sync.Once
	ch chan struct{}
}

func NewMyChan() *MyChan {
	return &MyChan{
		ch: make(chan struct{}),
	}
}

func (c *MyChan) Close() {
	c.Do(func() {
		close(c.ch)
	})
}

type MyconcurrentMap struct {
	sync.Mutex
	mp     map[int]int
	key2Ch map[int]*MyChan // 标记当前key 是否存在
}

func NewMyconcurrentMap() *MyconcurrentMap {
	return &MyconcurrentMap{
		mp:     make(map[int]int),
		key2Ch: make(map[int]*MyChan),
	}
}

func (m *MyconcurrentMap) Put(k, v int) {
	m.Lock()
	defer m.Unlock()
	m.mp[k] = v

	ch, ok := m.key2Ch[k]
	if !ok {
		return
	}
	ch.Close()

	// 1.
	// select {
	// case <-ch:
	// 	return
	// default:
	// 	close(ch)
	// }

	// 不能重复关闭close
	// close(ch) // close后唤醒全部阻塞等待的Goroutine
	// ch <- struct{}{}
}

func (m *MyconcurrentMap) Get(k int, maxWaitingDuration time.Duration) (int, error) {
	m.Lock()
	v, ok := m.mp[k]
	if ok {
		m.Unlock()
		return v, nil
	}

	// 不会被放多次 这个Channel会被复用
	ch, ok := m.key2Ch[k]
	if !ok {
		ch = NewMyChan()
		// ch = make(chan struct{})
		m.key2Ch[k] = ch
	}
	m.Unlock()

	tCtx, cancel := context.WithTimeout(context.Background(), maxWaitingDuration)
	defer cancel()

	select {
	case <-tCtx.Done():
		return -1, tCtx.Err()
	case <-ch.ch:
		// 直接跳过
	}
	m.Lock()
	v = m.mp[k]
	m.Unlock()

	// <-m.key2Ch[k]

	return v, nil
}

func main() {
	// Example usage
	cmap := NewMyconcurrentMap()

	go func() {
		time.Sleep(2 * time.Second)
		cmap.Put(1, 100)
		cmap.Put(2, 100)
	}()

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			val, err := cmap.Get(1, 5*time.Second)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(val)
		}()
	}

	val, err := cmap.Get(1, 1*time.Second)
	if err != nil {
		println("Get error:", err.Error())
	} else {
		println("Get value:", val)
	}
	wg.Wait()
}
