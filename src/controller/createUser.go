package controller

import (
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	err := rest_err.NewBadRequestError("Voce Chamou a rota de forma errada")
	c.JSON(err.Code, err)

}
