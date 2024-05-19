// package singleflight
package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"

	"golang.org/x/sync/singleflight"
)

var (
	sf           singleflight.Group
	ErrCacheMiss = errors.New("cache miss")
)

var wg sync.WaitGroup

func GetArticle(id int) (article string, err error) {
	// 假设对数据库进行调研，模拟不同并发下耗时不同
	// atomic.AddInt32(&count, 1)
	// time.Sleep(time.Duration(count))
	// sf.Do(key string, fn func() (interface{}, error))
	return "", nil
}

func SingleTest() {
	wg.Add(10)

	// mock 10 concurrency
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			data, err := Load("keysss")
			if err != nil {
				log.Println(err)
				return
			}
			log.Println(data)
		}()
	}
	wg.Wait()
}

// get data
func Load(key string) (string, error) {
	data, err := loadFromCache(key)
	if err != nil && err == ErrCacheMiss {
		// use singleflight
		v, err, _ := sf.Do(key, func() (interface{}, error) {
			data, err := loadFromDB(key)
			if err != nil {
				return nil, err
			}
			setCache(key, data)
			return data, nil
		})
		if err != nil {
			return "", err
		}
		data = v.(string)
	}
	return data, nil
}

// getDataFrom Cache. Mock cache miss
func loadFromCache(key string) (string, error) {
	return "", ErrCacheMiss
}

// setCache 写入缓存
func setCache(key, data string) {
}

// getDataFromDB
func loadFromDB(key string) (string, error) {
	fmt.Println("query db")
	unix := strconv.Itoa(int(time.Now().UnixNano()))
	return unix, nil
}

func main() {
	SingleTest()
	WZsingleFlight()
}
