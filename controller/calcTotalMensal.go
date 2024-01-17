package controller



import (
	"net/http" 
	"github.com/gin-gonic/gin"
"example.com/gogin/config"
	"example.com/gogin/models"

)
 
// @Summary Returns the specific unit with information for the pdt
// @Description Returns the unit
// @Accept  json
// @Produce  json
// @Param unidade body string true "Unidade"
// @Success 200 {array} models.PDT
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /unidade [get]
func TotCalcMensalFolha(context *gin.Context){
var totalMensalFolha float64= 0.0
var totalRhfolha int=0
	var pdts []models.PDT
	config.DB.Find(&pdts)
somaPortotalsalariomensalfolhaTotalUnidadeFuncaoPdt:= make(map[float64]map[string]map[string]float64)
somaPorRH:= make(map[string]map[string]int)


for _, pdt := range pdts {
	if somaPorRH[pdt.Unidade] == nil {
		somaPorRH[pdt.Unidade] = make(map[string]int)
	 }
	somaPorRH[pdt.Unidade][pdt.Funcao] +=1
	totalRhfolha+=1
}

	for _, pdt := range pdts {
		
		 if _, ok := somaPortotalsalariomensalfolhaTotalUnidadeFuncaoPdt[pdt.Total_mensal_folha]; !ok {
			somaPortotalsalariomensalfolhaTotalUnidadeFuncaoPdt[pdt.Total_mensal_folha] = make(map[string]map[string]float64)
		 }
		 if _, ok := somaPortotalsalariomensalfolhaTotalUnidadeFuncaoPdt[pdt.Total_mensal_folha][pdt.Unidade]; !ok {
			somaPortotalsalariomensalfolhaTotalUnidadeFuncaoPdt[pdt.Total_mensal_folha][pdt.Unidade] = make(map[string]float64)
		 }

	
		 somaPortotalsalariomensalfolhaTotalUnidadeFuncaoPdt[pdt.Total_mensal_folha][pdt.Unidade][pdt.Funcao]=0.0
		 somaPortotalsalariomensalfolhaTotalUnidadeFuncaoPdt[pdt.Total_mensal_folha][pdt.Unidade][pdt.Funcao] += pdt.Total_mensal_folha
	 }


	for _, pdt:= range pdts {
		
		totalMensalFolha+=somaPortotalsalariomensalfolhaTotalUnidadeFuncaoPdt[pdt.Total_mensal_folha][pdt.Unidade][pdt.Funcao] 
		
	}

	context.JSON(http.StatusOK, gin.H{
		"TotalMensalFolha" : totalMensalFolha,
		"TotalRhFolha" : totalRhfolha,
	})

	//somaPortotalsalariomensalfolhaTotalUnidadeFuncaoPdt[pdt.Total_mensal_folha][pdt.Unidade][pdt.Funcao]=0.0
    totalMensalFolha=0.0
	totalRhfolha=0
}