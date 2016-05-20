package controllers

import (
	"net/http"

    "godemo/createmodule/example/models"

	"github.com/gin-gonic/gin"
)

type TempsController struct{

	ApplicationController

}

func (ctrl TempsController) Index(c *gin.Context) {

	var temps []models.Temp
    ctrl.DB.Model(&models.Temp{}).Find(&temps)
    c.HTML(http.StatusOK, "index.html", gin.H{
    		"title" : "index",
    		"temps" : temps,
    });

}


