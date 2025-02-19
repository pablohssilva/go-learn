package main

import (
	"fmt"
	"log"
	"github.com/joho/godotenv"
)

func main(){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	routes.InitRoutes()
}

//router := gin.New() A diferença do new é que ele não instancia nenhun handler e nenhum middleware!
router := gin.Default()

routes.InitRoutes(&router.RouterGroup)

if err := router.Run(":8080"); err != nil {
	log.Fatal(err)
}
