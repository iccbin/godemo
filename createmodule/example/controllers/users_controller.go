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
	var users []models.User
	err := ctrl.DB.Model(&models.User{}).Find(&users).Error

	if err != nil {
		ctrl.Render.Text(c.Writer, http.StatusOK, err.Error())
	} else {
		ctrl.Render.HTML(c.Writer, http.StatusOK, "users/index", gin.H{
			"title": "index",
			"users": users,
		})
	}

}

func (ctrl UsersController) New(c *gin.Context) {
	ctrl.Render.HTML(c.Writer, http.StatusOK, "users/new", gin.H{
		"title": "new",
	})
}

func (ctrl UsersController) Create(c *gin.Context) {
	money,err := strconv.ParseFloat(c.PostForm("money"), 64)
	if err != nil {
		ctrl.Render.Text(c.Writer, http.StatusOK, err.Error())
		return
	}
	name := c.PostForm("name")
	if name == "" {
		ctrl.Render.Text(c.Writer, http.StatusOK, "name is not null")
		return
	}
	password := c.PostForm("password")
	gender,err := strconv.ParseBool(c.PostForm("gender"))
	if err != nil {
		ctrl.Render.Text(c.Writer, http.StatusOK, err.Error())
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
		ctrl.Render.Text(c.Writer, http.StatusOK,  err.Error())
	} else {
		ctrl.Render.Text(c.Writer, http.StatusOK, "success")
	}

}

func (ctrl UsersController) Delete(c *gin.Context) {
	//user := models.User{}
	//err := ctrl.DB.Delete(&user, c.Param("id")).Error
	//if err != nil {
	//	ctrl.Render.Text(c.Writer, http.StatusOK, err.Error())
	//} else {
	//	ctrl.Render.Text(c.Writer, http.StatusOK, "success")
	//}
	ctrl.Render.Text(c.Writer, http.StatusOK, "success")
}

func (ctrl UsersController) Edit(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))

	user := models.User{ID: uint(userId)}

	err := ctrl.DB.First(&user).Error
	if err != nil {
		ctrl.Render.Text(c.Writer, http.StatusOK, err.Error())
	} else {
		ctrl.Render.HTML(c.Writer, http.StatusOK, "users/edit", gin.H{
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
		ctrl.Render.Text(c.Writer, http.StatusOK, err.Error())
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
		ctrl.Render.Text(c.Writer, http.StatusOK, err.Error())
	} else {
		ctrl.Render.Text(c.Writer, http.StatusOK, "success")
	}

}

func (ctrl UsersController) Show(c *gin.Context)  {
	var user models.User
	err := ctrl.DB.First(&user, c.Param("id")).Error

	if err != nil {
		ctrl.Render.Text(c.Writer, http.StatusOK, err.Error())
	} else {
		ctrl.Render.HTML(c.Writer,http.StatusOK,"users/show",gin.H{
			"title":"show",
			"user":user,
		})
	}
}
