package controllers

import (
	"github.com/jinzhu/gorm"

)

type ApplicationController struct {
	DB *gorm.DB
}

