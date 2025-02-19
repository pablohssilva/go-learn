package controller

import "github.com/gin-gonic/gin"

func CreateUser(c *gin.Context){
>>>>>>> c27e590 (referencia)

func CreateUser(c *gin.Context) {

	err := rest_err.NewBadRequestError("")
	c.JSON(err.Code, err)
	
}
>>>>>>> c27e590 (referencia)
