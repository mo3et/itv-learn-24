package main

import (
	"log"

	"github.com/gin-gonic/gin"
	example "github.com/mo3et/itv-learn-24/gin-demo/example/single-example"
)

func main() {
	// create default router engine.
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		// c.JSON: 返回JSON格式的数据
		c.JSON(200, gin.H{
			"message": "hello,world!",
		})
	})
	// example setup
	example.SetupBookRoutes(r)
	example.SetupHTMLRender(r)

	// example.SetCustomTemplateRender(r)
	example.SetFormatAsDateFunc(r)

	// start HTTP service.
	if err := r.Run(":8080"); err != nil {
		log.Fatal("router run failed.")
	}
}
