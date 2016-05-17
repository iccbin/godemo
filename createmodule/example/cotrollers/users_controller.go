package cotrollers

import (
	"net/http"
	"fmt"

	"godemo/createmodule/example/models"

	"github.com/gin-gonic/gin"
)

type UsersController struct{
	ApplicationController
}


func (ctrl UsersController) Index(c *gin.Context) {
	var users []models.User
	ctrl.DB.Model(&models.User{}).Find(&users)
	fmt.Println(users)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title" : "index",
	});
}

//func (user *UsersController) Show(c *gin.Context) {
//	c.HTML(http.StatusOK, "show.html", gin.H{
//		"title" : "show",
//	})
//}
//
//func (user *UsersController) Create(c *gin.Context) {
//	c.HTML(http.StatusOK, "show.html", gin.H{
//		"title" : "show",
//	})
//}
//
//func (user *UsersController) Delete(c *gin.Context) {
//	c.HTML(http.StatusOK, "show.html", gin.H{
//		"title" : "show",
//	})
//}
//
//func (user *UsersController) Update(c *gin.Context) {
//	c.HTML(http.StatusOK, "show.html", gin.H{
//		"title" : "show",
//	})
//}