package controller

import (
	uc "github.com/georgexx009-wizeline/ondemand-go-bootcamp/use-cases"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	Router   *gin.Engine
	UseCases *uc.UseCases
}

func NewController(u *uc.UseCases) *Controller {
	return &Controller{UseCases: u}
}
