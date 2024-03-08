package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mo3et/itv-learn-24/gin-demo/example/gin-DB/util"
)

func Router(r *gin.Engine) {
	util.NewMysqlDB()

	s := r.Group("/student")
	{
		s.PUT("/save", api.Save)
	}
}
