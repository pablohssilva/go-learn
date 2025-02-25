package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pablohssilva/go-learn/src/configuration/rest_err"
)

func CreateUser(c *gin.Context) {

	err := rest_err.NewBadRequestError("Voce Chamou a rota de forma errada")
	c.JSON(err.Code, err)

}
