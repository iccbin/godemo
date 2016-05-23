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
        c.String(http.StatusOK, err.Error())
    } else {
        c.HTML(http.StatusOK, "index.html", gin.H{
                "title": "index",
                "users": users,
        })
    }

}

func (ctrl UsersController) New(c *gin.Context) {
	c.HTML(http.StatusOK, "new.html", gin.H{
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
    		c.String(http.StatusOK, err.Error())
    		return
    }
	user := models.User{
	            Age:  age,
	            Gender:  gender,
	            Money:  money,
	            Name:  name,
	            Password:  password,
	        }
	err = ctrl.DB.Create(&user).Error
    if err != nil {
    		c.String(http.StatusOK, "create user error")
    } else {
    		c.Redirect(http.StatusMovedPermanently, "/")
    }
}

func (ctrl UsersController) Delete(c *gin.Context) {
	userId, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
        c.String(http.StatusOK, err.Error())
        return
    }

	user := models.User{ID: userId}

	err := ctrl.DB.Delete(&user).Error
	if err != nil {
		c.String(http.StatusOK, err.Error())
	} else {
		c.String(http.StatusOK, "ok")
	}
}

func (ctrl UsersController) Edit(c *gin.Context) {
	userId, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
            c.String(http.StatusOK, err.Error())
            return
    }

	user := models.User{ID: userId}

	err = ctrl.DB.First(&user).Error
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
	
    age, err := strconv.ParseInt(c.PostForm("age"),10,0)
    gender, err := strconv.ParseBool(c.PostForm("gender"))
    id, err := strconv.ParseUint(c.PostForm("id"),10,0)
    money, err := strconv.ParseFloat(c.PostForm("money"), 64)
    name := c.PostForm("name")
    password := c.PostForm("password")
    if err != nil {
        c.String(http.StatusOK, err.Error())
        return
    }

	user := models.User{
                    Age:  age,
                    Gender:  gender,
                    ID:  id,
                    Money:  money,
                    Name:  name,
                    Password:  password,
            }
	err := ctrl.DB.Save(&user).Error
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
         c.HTML(http.StatusOK,"show.html",gin.H{
    		"title":"show",
    		"user":user,
    	 })
    }


}