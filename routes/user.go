package routes

import (
	//"net/http"
	"github.com/gin-gonic/gin"
	"example.com/gogin/controller"
)
/*
func Rotasadicionalnoturno(router *gin.Engine) {
	router.GET("/adicional_noturno", controller.Adicionalnoturno) // Corrigi o nome da função
	//router.GET("/auxilio_alimentacao, ")
	//router.GET("/salariobase", controller.Salariobasee)

	//corrigir

}
*/
/*func Rotaauxilioalimentacao(router *gin.Engine){
	router.GET("/auxilio_alimentacao",controller.Auxilioalimentac)

}

func DeducaoInss(router *gin.Engine){

	router.GET("/deducao_inss", controller.Deducaoinsss)

}
*/
/*
func Gratificacao(router *gin.Engine){

	router.GET("/gratificacao", controller.Gratificaoo)
}

func Horaextra(router *gin.Engine){
	router.GET("/horaextra",controller.Horaextraa)
}

func Insalubridade(router *gin.Engine){
	router.GET("/insalubridade", controller.Insalubridadee)
}

func Periculosidade(router *gin.Engine){
	router.GET("/periculosidade", controller.Periculosidadee)
}

/*func Salariobase(router *gin.Engine){
	router.GET("/salariobase", controller.Salariobasee)
} */


func Demos(router *gin.Engine){
	router.GET("/tudo",controller.GetUserss)
}

func Uni(router *gin.Engine){
	router.POST("/unidade",controller.GetUserss)
}

func Calc(router *gin.Engine){
	router.GET("/opa", controller.Calculo)
}



//func Getiadvh(router *gin.Engine){
//	router.GET("/iadvh",controller.Getiadvh)
//}
/*
func Getiadvhs(router *gin.Engine){
	router.GET("/rh",controller.Getiadvhh)
}

func Getiadvhss(router *gin.Engine){
	router.GET("/rht",controller.Getiadvhhh)
}

func Getiadvhssss(router *gin.Engine){
	router.GET("/rhtt",controller.Getiadvhhhh)
}

func Soma(router *gin.Engine){
	router.GET("/somarh",controller.Somarh)
}

func Somadois(router *gin.Engine){
	router.GET("/somadois",controller.Somarhdois)
}

func Somainsalubridade(router *gin.Engine){
	router.GET("/insalubridade",controller.Somainsalubridade)
}
func Somadeducao(router *gin.Engine){
	router.GET("/deducao",controller.Somadeducaoinss)
}

func Somaencargos(router *gin.Engine){
	router.GET("/encargos",controller.Somaencargos)
}

func Somasalariobase(router *gin.Engine){
	router.GET("salariobase",controller.Somadesalariobase)
}


func Somahoraextra(router *gin.Engine){
	router.GET("horaextra",controller.Somahoraextra)
}

func Somagratificacao(router *gin.Engine){
	router.GET("gratificacao",controller.Somagratificacao)
}

func Somapericulosidade(router *gin.Engine){
	router.GET("periculosidade",controller.Somapericulosidade)
}

*/
//func Somaauxilioalimentacao(router *gin.Engine){
//	router.GET("auxilio",controller.Somaauxilioalimentacao)
//}
/*
func Somaadicionalnoturno(router *gin.Engine){
	router.GET("adicional",controller.Somaadicionalnoturno)
}

*/


/*func Getiadvhsss(router *gin.Engine){
	router.GET("/insalubridade",controller.Insalubridade)
}

*/
