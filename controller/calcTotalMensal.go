package controller



import (
	"net/http" 
	"github.com/gin-gonic/gin"
"example.com/gogin/config"
	"example.com/gogin/models"

)
 

func TotCalcMensalFolha(context *gin.Context){
var totalMensalFolha float64= 0.0
	var pdts []models.PDT
	config.DB.Find(&pdts)
	somaPortotalsalariomensalfolhaTotalUnidadeFuncaoPdt:= make(map[float64]map[string]map[string]float64)

		/*
		
	for _, pdt := range pdts {
	   /*
		if _, ok := somaPortotalsalariomensalfolhaTotalUnidadeFuncaoPdt[pdt.Total_mensal_folha]; !ok {
			somaPortotalsalariomensalfolhaTotalUnidadeFuncaoPdt[pdt.Total_mensal_folha] = make(map[string]float64)
		}
		if _, ok := somaPortotalsalariomensalfolhaTotalUnidadeFuncaoPdt[pdt.Total_mensal_folha][pdt.Unidade]; !ok {
			somaPortotalsalariomensalfolhaTotalUnidadeFuncaoPdt[pdt.Total_mensal_folha]=make(map[string]float64)
		}

		//soma1:=somaPorencargosUnidadeFuncaoPdt 
		//soma3:= somaPortotalsalariofolhaUnidadeFuncaoPdt
		somaPortotalsalariomensalfolhaTotalUnidadeFuncaoPdt[pdt.Total_mensal_folha]+= pdt.Total_mensal_folha
		//totalMensalFolha=somaPortotalsalariomensalfolhaTotalUnidadeFuncaoPdt[pdt.Total_mensal_folha]
	}

*/
	for _, pdt := range pdts {
		// soma2=0.0
		 if _, ok := somaPortotalsalariomensalfolhaTotalUnidadeFuncaoPdt[pdt.Total_mensal_folha]; !ok {
			somaPortotalsalariomensalfolhaTotalUnidadeFuncaoPdt[pdt.Total_mensal_folha] = make(map[string]map[string]float64)
		 }
		 if _, ok := somaPortotalsalariomensalfolhaTotalUnidadeFuncaoPdt[pdt.Total_mensal_folha][pdt.Unidade]; !ok {
			somaPortotalsalariomensalfolhaTotalUnidadeFuncaoPdt[pdt.Total_mensal_folha][pdt.Unidade] = make(map[string]float64)
		 }

		 //soma1:=somaPorencargosUnidadeFuncaoPdt 
		 //soma3:= somaPortotalsalariofolhaUnidadeFuncaoPdt
		 somaPortotalsalariomensalfolhaTotalUnidadeFuncaoPdt[pdt.Total_mensal_folha][pdt.Unidade][pdt.Funcao]=0.0
		 somaPortotalsalariomensalfolhaTotalUnidadeFuncaoPdt[pdt.Total_mensal_folha][pdt.Unidade][pdt.Funcao] += pdt.Total_mensal_folha
	 }


	for _, pdt:= range pdts {
		
		totalMensalFolha+=somaPortotalsalariomensalfolhaTotalUnidadeFuncaoPdt[pdt.Total_mensal_folha][pdt.Unidade][pdt.Funcao]

	}

	context.JSON(http.StatusOK, gin.H{
		"TotalMensalFolha" : totalMensalFolha,
	})



}