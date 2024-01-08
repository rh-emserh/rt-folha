package main 


import (
	//"net/http"

	"github.com/gin-gonic/gin"
	"example.com/gogin/routes"
	"example.com/gogin/config"
	"example.com/gogin/models"
)


// @title 

func main(){

	router:= gin.New()
	config.Connect()
	config.DB.AutoMigrate(&models.PDT{})
    config.DB.AutoMigrate(&models.Iadvh{})
	routes.Demos(router)
	routes.Uni(router)	

	//routes.Somaauxilioalimentacao(router)
	router.Run("localhost:8080")

}