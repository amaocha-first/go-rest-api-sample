package main

import (
	"rest-api-sample/article"
	"rest-api-sample/httpd/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	article := article.New()

	r := gin.Default()
	r.GET("/article", handler.ArticlesGet(article))
	r.POST("/article", handler.ArticlePost(article))
	r.Run()
}
