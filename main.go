package main 


import (
	//"net/http"

	"github.com/gin-gonic/gin"
	"example.com/gogin/routes"
	"example.com/gogin/config"
	//"example.com/gogin/models"
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
	router.Use(cors.Default())
	routes.Uni(router)	
	
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	router.Use(cors.New(config))


	routes.SwaggerUni(router)
	routes.TotalMensalFolhaPagamento(router)
	routes.Totalsalariobase(router)
	routes.AdicionalNoturno(router)


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