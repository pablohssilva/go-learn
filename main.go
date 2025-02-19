package main

import (
	
	"log"
	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
)

func main(){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	//router := gin.New() A diferença do new é que ele não instancia nenhun handler e nenhum middleware!
	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup)

	if err := router.Run(":8080"); err != nil {
	log.Fatal(err)
	}
}