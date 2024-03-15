package main

import (
	"github.com/gin-gonic/gin"
	routers "github.com/mo3et/itv-learn-24/gin-demo/example/gin-DB/router"
)

func main() {
	r := gin.Default()
	routers.Router(r)
	if err := r.Run(":8001"); err != nil {
		panic(err)
	}
}
