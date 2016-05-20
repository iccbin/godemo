package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

)

type UsersController struct {
	ApplicationController
}

func (ctrl UsersController) Index(c *gin.Context) {
	var users []models.User
	ctrl.DB.Model(&models.User{}).Find(&users)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "index",
		"users": users,
	})
}

func (ctrl UsersController) New(c *gin.Context) {
	c.HTML(http.StatusOK, "new.html", gin.H{
		"title": "new",
	})
}

func (ctrl UsersController) Create(c *gin.Context) {
	user := models.User{
	            Name:  c.PostForm("name"),
	            Password:  c.PostForm("password"),
	        }
	ctrl.DB.NewRecord(user)
	ctrl.DB.Create(&user)

	c.Redirect(http.StatusMovedPermanently, "/")
}

func (ctrl UsersController) Delete(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))
	user := models.User{ID: uint(userId)}

	err := ctrl.DB.Delete(&user).Error
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
	userId, _ := strconv.Atoi(c.Param("id"))
	id := uint(userId)  
	name := c.PostForm("name") 
	password := c.PostForm("password")
	user := models.User{
                ID:  c.PostForm("id"),
                Name:  c.PostForm("name"),
                Password:  c.PostForm("password"),
            }
	err := ctrl.DB.Save(&user).Error
	if err != nil {
		c.String(http.StatusOK, err.Error())
	} else {
		c.String(http.StatusOK, "ok")
	}

}

func (ctrl UsersController) Show(c *gin.Context)  {
	var users []models.User
	ctrl.DB.Where("id > ?", c.Param("key")).Find(&users)

	c.HTML(http.StatusOK,"show.html",gin.H{
		"title":"hello",
		"users":users,
	})

}