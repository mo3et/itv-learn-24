package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/singleflight"
)

func getData(id int64) string {
	fmt.Println("query...")
	time.Sleep(10 * time.Second)
	return "liwenzhou.com"
}

var sg singleflight.Group

func WZsingleFlight() {
	g := new(singleflight.Group)

	// first call
	go func() {
		v1, _, shared := g.Do("getData", func() (interface{}, error) {
			ret := getData(1)
			return ret, nil
		})
		fmt.Printf("1st call: v1:%v, shared:%v\n", v1, shared)
	}()

	time.Sleep(2 * time.Second)

	// 2nd call
	v2, _, shared := g.Do("getData", func() (interface{}, error) {
		ret := getData(1)
		return ret, nil
	})
	fmt.Printf("2nd call: v2:%v, shared:%v\n", v2, shared)
	fmt.Println(g)
}

func doChanGetData(ctx context.Context, g *singleflight.Group, id int64) (string, error) {
	ch := g.DoChan("getChanData", func() (interface{}, error) {
		res := getData(id)
		return res, nil
	})
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-ch:
		return res.Val.(string), res.Err
	}
}
