package main 


import (
	//"net/http"

	"github.com/gin-gonic/gin"
	"example.com/gogin/routes"
	"example.com/gogin/config"
	"example.com/gogin/models"
	_ "example.com/gogin/docs"
	"log"
	"os"
	"github.com/gin-contrib/cors"
)


// @title 
// @title API FOLHA DE PAGAMENTO
// @version 1.0
// @description API for crud operations on users
// @BasePath /
func main(){

	router:= gin.New()
	config.Connect()
	config.DB.AutoMigrate(&models.PDT{})
    config.DB.AutoMigrate(&models.Iadvh{})
	router.Use(cors.Default())
	//routes.Demos(router)
	routes.Uni(router)	
	
	//routes.Somaauxilioalimentacao(router)
	//router.Run("localhost:8080")

	routes.SwaggerUni(router)
	routes.TotalMensalFolhaPagamento(router)

	// Adicionando o bloco para configurar o host e a porta
	host := os.Getenv("HOST")
	if host == "" {
		host = "0.0.0.0"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	address := host + ":" + port
	if err := router.Run(address); err != nil {
		log.Panicf("error: %s", err)
	}

}