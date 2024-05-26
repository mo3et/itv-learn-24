package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		t := time.NewTicker(time.Second * 1)
		for {
			select {
			case <-t.C:
				go func() {
					defer func() {
						if err := recover(); err != nil {
							fmt.Println(err)
						}
					}()
					proc()
				}()
			}
		}
	}()

	select {
		
	}
}

func proc() {
	panic("ok")
}
