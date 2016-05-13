package main

import (

	"github.com/gin-gonic/gin"
	"godemo/restful/controllers"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	//router.GET("/", func(c *gin.Context) {c.String(http.StatusOK,"ttt")})
	router.GET("/", controllers.Index)
	router.POST("/articles", controllers.ArticleAdd)
	router.DELETE("/articles/:id", controllers.ArticleDelete)
	router.GET("/articles/:id", controllers.ArticleUpdate)
	router.GET("/articles/:id/edit", controllers.ArticleUpdate)
	router.PUT("/articles/:id",controllers.ArticleUpdate)
	////添加文章
	//router.POST("/article/add", controllers.ArticleAdd)
	////删除文章
	////router.GET("/article/delete/:id", controllers.ArticleDel)
	////查看文章
	//router.GET("/article/:id", controllers.ArticleUp)
	////修改文章
	//router.POST("/article/update/:id", controllers.ArticleUp)
	////查找文章
	//router.POST("/article/find", controllers.ArticleFind)

	router.Run(":9000")
}
