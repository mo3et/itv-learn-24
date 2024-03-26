package main

import (
	"fmt"
	"sync"
	"time"
)

type Repo interface {
	Get() (string, error)
	Set() error
}

type repoImpl struct {
	// 刚好需要锁时， 就用这个var
	sync.Mutex
}

func (r *repoImpl) Get() (string, error) {
	r.Lock()
	defer r.Unlock()
	time.Sleep(10 * time.Second)
	return "", nil
}

func (r *repoImpl) Set() error {
	r.Lock()
	defer r.Unlock()
	time.Sleep(3 * time.Second)
	return nil
}

func NewRepo() Repo {
	return &repoImpl{}
}

func main() {
	rr := NewRepo()
	start := time.Now() // 记录开始时间
	defer func() {
		fmt.Printf("运行完成，耗时：%s\n", time.Since(start))
	}()

	go func() {
		if err := rr.Set(); err != nil {
			fmt.Println(err)
		}
	}()
	if _, err := rr.Get(); err != nil {
		fmt.Println(err)
		return
	}
}
