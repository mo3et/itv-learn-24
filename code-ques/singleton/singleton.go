package main

import (
	"fmt"
	"sync"
)

var (
	mu             = &sync.Mutex{}
	singleInstance *single
	once           sync.Once
)

type single struct{}

func getInstance() *single {
	if singleInstance == nil {
		mu.Lock()
		defer mu.Unlock()
		if singleInstance == nil {
			fmt.Println("creating single instance now.")
			singleInstance = &single{}
		} else {
			fmt.Println("Single instance already created.")
		}
	}
	return singleInstance
}

func main() {
	for i := 0; i < 30; i++ {
		go getInstance()
		go getInstanceOnce(i)
	}

	// Scanln is similar to Scan, but stops scanning at a newline and
	// after the final item there must be a newline or EOF.
	fmt.Scanln()
}

func getInstanceOnce(i int) *single {
	if singleInstance == nil {
		once.Do(
			func() {
				fmt.Println("Creating single instance now.", i)
				singleInstance = &single{}
			})
	} else {
		fmt.Println("Single instance already created. ", i)
	}
	return singleInstance
}
