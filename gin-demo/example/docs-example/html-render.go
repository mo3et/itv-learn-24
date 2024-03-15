package example

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func SetupHTMLRender(r *gin.Engine) {
	r.LoadHTMLGlob("templates/**/*")
	// r.LoadHTMLFiles("templates/posts/index.tmpl", "templates/users/index.tmpl")
	r.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "post/index.tmpl", gin.H{
			"title": "posts/index",
		})
	})

	r.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
			"title": "users/index",
		})
	})
}

// func SetCustomTemplateRender(r *gin.Engine) {
// 	// r.LoadHTMLGlob("templates/**/*")
// 	html := template.Must(template.ParseFiles("users/index.tmpl", "posts/index.tmpl"))
// 	r.SetHTMLTemplate(html)
// }

func SetCustomDelimiter(r *gin.Engine) {
	r.Delims("{[{", "}]}")
	// r.LoadHTMLGlob("/path/to/templates")
}

func SetFormatAsDateFunc(r *gin.Engine) {
	SetCustomDelimiter(r)
	r.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})
	r.LoadHTMLFiles("./templates/raw.tmpl")

	r.GET("/raw", func(c *gin.Context) {
		c.HTML(http.StatusOK, "raw.tmpl", map[string]interface{}{
			"now": time.Date(2024, 0o3, 8, 3, 2, 0, 8, time.UTC),
		})
		// params:=c.Query("params")
		c.JSON(http.StatusOK, gin.H{
			"message":        "raw",
			"counter-strike": "csgo",
			"krystal":        3208,
		})
	})
}

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d/%02d/%02d", year, month, day)
}

func SetStaticFileRender(r *gin.Engine) {
	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/**/*")
}
