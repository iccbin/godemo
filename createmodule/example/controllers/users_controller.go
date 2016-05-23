package controllers

import (
	"net/http"
	"strconv"

	"godemo/createmodule/example/models"

	"github.com/gin-gonic/gin"

)

type UsersController struct {
	ApplicationController
}

func (ctrl UsersController) Index(c *gin.Context) {
	//var users []models.User
	//err := ctrl.DB.Model(&models.User{}).Find(&users).Error
	//if err != nil {
	//	c.String(http.StatusOK, err.Error())
	//} else {
	//	c.HTML(http.StatusOK, "index.html", gin.H{
	//		"title": "index",
	//		"users": users,
	//	})
	//}

	ctrl.Render.HTML(c.Writer, http.StatusOK, "users/show", gin.H{
		"title":"hello",
	})

}

func (ctrl UsersController) New(c *gin.Context) {
	c.HTML(http.StatusOK, "new.html", gin.H{
		"title": "new",
	})
}

func (ctrl UsersController) Create(c *gin.Context) {
	money,err := strconv.ParseFloat(c.PostForm("money"), 64)
	name := c.PostForm("name")
	password := c.PostForm("password")
	gender,err := strconv.ParseBool(c.PostForm("gender"))
	if err != nil {
		c.String(http.StatusOK, err.Error())
		return
	}
	user := models.User{
		Name: name,
		Password: password,
		Money:money,
		Gender:gender,
	}
	err = ctrl.DB.Create(&user).Error
	if err != nil {
		c.String(http.StatusOK, "create user error")
	} else {
		c.Redirect(http.StatusMovedPermanently, "/")
	}

}

func (ctrl UsersController) Delete(c *gin.Context) {
	user := models.User{}
	err := ctrl.DB.Delete(&user, c.Param("id")).Error
	if err != nil {
		c.String(http.StatusOK, err.Error())
	} else {
		c.String(http.StatusOK, "ok")
	}
}

func (ctrl UsersController) Edit(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))

	user := models.User{ID: uint(userId)}

	err := ctrl.DB.First(&user).Error
	if err != nil {
		c.String(http.StatusOK, err.Error())
	} else {
		c.HTML(http.StatusOK, "edit.html", gin.H{
			"title": "edit",
			"user":  user,
		})
	}
}

func (ctrl UsersController) Update(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	name := c.PostForm("name")
	password := c.PostForm("password")
	money, err := strconv.ParseFloat(c.PostForm("money"), 64)
	gender, err := strconv.ParseBool(c.PostForm("gender"))
	if err != nil {
		c.String(http.StatusOK, err.Error())
		return
	}
	user := models.User{
		ID: uint(userId),
		Name: name,
		Password: password,
		Money:money,
		Gender:gender,
	}

	err = ctrl.DB.Save(&user).Error
	if err != nil {
		c.String(http.StatusOK, err.Error())
	} else {
		c.String(http.StatusOK, "ok")
	}

}

func (ctrl UsersController) Show(c *gin.Context)  {
	var user models.User
	err := ctrl.DB.First(&user, c.Param("id")).Error

	if err != nil {
		c.String(http.StatusOK, err.Error())
	} else {
		//w http.ResponseWriter, status int, name string, binding interface{}, htmlOpt ...HTMLOptions
		ctrl.Render.HTML(c.Writer, http.StatusOK,"show.html",gin.H{
			"title":"hello",
			"user":user,
		})
	}
}
