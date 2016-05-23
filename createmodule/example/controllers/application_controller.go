package controllers

import (
	"github.com/jinzhu/gorm"

	"github.com/unrolled/render"
)

type ApplicationController struct {
	DB *gorm.DB
	Render *render.Render
}


