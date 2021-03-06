package controllers

import (

	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"

	"godemo/restful/models"
)

func Index(c *gin.Context){

	articles,_ := models.AllArticles()

	c.HTML(http.StatusOK,"index.html",gin.H{
		"title":"aggghhhhg",
		"articles":articles,
	})
}

func ArticleAdd(c *gin.Context)  {
	article := models.Article{Title:c.PostForm("title"),
		Content:c.PostForm("content")}
	err := article.Insert()
	if err!=nil {
		c.String(http.StatusOK, err.Error())
	} else {
		c.Redirect(http.StatusMovedPermanently,"/")
	}
}

func ArticleDelete(c *gin.Context)  {
	articleId,_ := strconv.Atoi(c.Param("id"))

	article := models.Article{Id:articleId}
	err := article.Delete()
	if err!=nil {
		c.String(http.StatusOK, err.Error())
	} else {
		c.String(http.StatusOK, "ok")
	}

}

func ArticleUpdate(c *gin.Context){
	articleId,_ := strconv.Atoi(c.Param("id"))
	article := models.Article{Id:articleId}
	title := c.PostForm("title")
	content := c.PostForm("content")
	article.Title = title;
	article.Content = content;
	article.Update()
	c.String(http.StatusOK, "ok")

}

func ArticleGetEdit(c *gin.Context){
	articleId,_ := strconv.Atoi(c.Param("id"))
	article := models.Article{Id:articleId}

	article.Find()
	c.HTML(http.StatusOK,"update.html",gin.H{
		"title":"hello",
		"article":article,
	})
}

func ArticleFind(c *gin.Context)  {
	articles,_ := models.FindArticles(c.PostForm("key"))

	c.HTML(http.StatusOK,"find.html",gin.H{
		"title":"hello",
		"articles":articles,
	})
}


