package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pablohssilva/go-learn/src/configuration/rest_err"
)
<<<<<<< HEAD
=======

func CreateUser(c *gin.Context){
>>>>>>> c27e590 (referencia)

func CreateUser(c *gin.Context) {

	err := rest_err.NewBadRequestError("")
	c.JSON(err.Code, err)

<<<<<<< HEAD
}
=======
}
>>>>>>> c27e590 (referencia)
