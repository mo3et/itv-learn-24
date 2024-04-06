package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	instance *Singleton
	mu       sync.Mutex
	once     sync.Once
	wg       sync.WaitGroup
)

type Singleton struct{}

func GetSingleton(i int) *Singleton {
	// fmt.Println(instance)
	fmt.Printf("%d is %v\n", i, instance)
	if instance == nil {
		mu.Lock()
		defer mu.Unlock()
		if instance == nil {
			fmt.Println("init Singleton.")
			instance = &Singleton{}
		}
	}
	return instance
}

func GetSingletonOnce() *Singleton {
	once.Do(func() {
		fmt.Println("init Singleton.")
		instance = &Singleton{}
	})
	return instance
}

func Haloa() func() int {
	s := 0
	return func() int {
		s++
		return s
	}
}

func main() {
	// ha := Haloa()
	// ha()
	// ha()
	// ha()
	// fmt.Println()

	// // s1 := make([]int, 5)
	// s1 := []int{1, 2, 3, 4, 5}
	// s2 := s1[:3] // 1,2,3
	// // s3 := make([]int, 3, 3)
	// s3 := s1[0:3:3]
	// s3 = s1[:3]
	// s2 = append(s2, 9)
	// s3 = append(s3, 9)
	// fmt.Println(s3)
	// fmt.Println(s2, &s2)
	// fmt.Println(s1, &(s1))

	// s := new(int)
	// a := make([]int, 3)
	// fmt.Println(s)
	// fmt.Println(a)

	startTime := time.Now()
	sumgo := 0
	for i := 0; i < 2000000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sumgo++
			GetSingleton(i)
		}()
	}

	wg.Wait()
	func() {
		elapsedTime := time.Since(startTime)
		fmt.Println("done when use ", elapsedTime)
		fmt.Println("sum goroutine is ", sumgo)
	}()
	// select {}
}
