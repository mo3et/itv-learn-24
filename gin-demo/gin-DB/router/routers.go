package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/mo3et/itv-learn-24/gin-demo/example/gin-DB/app/api"
	"github.com/mo3et/itv-learn-24/gin-demo/example/gin-DB/data"
)

func Router(r *gin.Engine) {
	// start DB
	data.Init()

	s := r.Group("/student")
	{
		s.PUT("/save", api.Save)
		s.GET("/select", api.SelectById)
		s.POST("/update", api.Update)
		s.DELETE("/delete", api.Delete)
	}
}
