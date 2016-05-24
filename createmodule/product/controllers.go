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
	err := ctrl.DB.Model(&models.User{}).Find(&users)

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
	ctrl.Render.HTML(c.Writer,http.StatusOK, "users/new", gin.H{
		"title": "new",
	})
}

func (ctrl UsersController) Create(c *gin.Context) {
    
    age, err := strconv.ParseInt(c.PostForm("age"),10,0)
    gender, err := strconv.ParseBool(c.PostForm("gender"))
    money, err := strconv.ParseFloat(c.PostForm("money"), 64)
    name := c.PostForm("name")
    password := c.PostForm("password")

    if err != nil {
    	ctrl.Render.Text(c.Writer, http.StatusOK, err.Error())
    	return
    }
	user := models.User{
	            Age:  age,
	            CreatedAt:  createdat,
	            DeletedAt:  deletedat,
	            Gender:  gender,
	            Money:  money,
	            Name:  name,
	            Password:  password,
	            UpdatedAt:  updatedat,
	        }
	err = ctrl.DB.Create(&user).Error
    if err != nil {
    		ctrl.Render.Text(c.Writer, http.StatusOK, "create user error")
    } else {
    		ctrl.Render.Text(c.Writer, http.StatusOK, "success")
    }
}

func (ctrl UsersController) Delete(c *gin.Context) {
	userId, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
        ctrl.Render.Text(c.Writer, http.StatusOK, err.Error())
        return
    }

	user := models.User{ID: userId}

	err := ctrl.DB.Delete(&user).Error
	if err != nil {
		ctrl.Render.Text(c.Writer, http.StatusOK, err.Error())
	} else {
		ctrl.Render.Text(c.Writer, http.StatusOK, "success")
	}
}

func (ctrl UsersController) Edit(c *gin.Context) {
	userId, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
            ctrl.Render.Text(c.Writer, http.StatusOK, err.Error())
            return
    }

	user := models.User{ID: userId}

	err = ctrl.DB.First(&user).Error
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
	
    age, err := strconv.ParseInt(c.PostForm("age"),10,0)
    gender, err := strconv.ParseBool(c.PostForm("gender"))
    id, err := strconv.ParseUint(c.PostForm("id"),10,0)
    money, err := strconv.ParseFloat(c.PostForm("money"), 64)
    name := c.PostForm("name")
    password := c.PostForm("password")
    if err != nil {
        ctrl.Render.Text(c.Writer, http.StatusOK, err.Error())
        return
    }

	user := models.User{
                    Age:  age,
                    CreatedAt:  createdat,
                    DeletedAt:  deletedat,
                    Gender:  gender,
                    ID:  id,
                    Money:  money,
                    Name:  name,
                    Password:  password,
                    UpdatedAt:  updatedat,
            }
	err := ctrl.DB.Save(&user).Error
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
        ctrl.Render.Text(c.Writer, "users/show",gin.H{
    		"title":"show",
    		"user":user,
    	 })
    }

}}