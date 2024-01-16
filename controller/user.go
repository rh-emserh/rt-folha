package controller

import (
	"net/http" // Uncomment this line
	"github.com/gin-gonic/gin"
	"example.com/gogin/config"
	"example.com/gogin/models"

//"gorm.io/gorm"

//"encoding/json"

"sync"

)


func Calculo() {
    //var pdt []models.PDT
	var pdts []models.PDT
	config.DB.Find(&pdts)

	var iadvhs []models.Iadvh
	config.DB.Find(&iadvhs)


            somaPortotalsalariomensalfolhaUnidadeFuncaoPdt := make(map[float64]map[string]map[string]float64)
		
		
			for _, pdt := range pdts {
            
				if _, ok := somaPortotalsalariomensalfolhaUnidadeFuncaoPdt[pdt.Total_mensal_folha]; !ok {
					somaPortotalsalariomensalfolhaUnidadeFuncaoPdt[pdt.Total_mensal_folha] = make(map[string]map[string]float64)
				}
				if _, ok := somaPortotalsalariomensalfolhaUnidadeFuncaoPdt[pdt.Total_mensal_folha][pdt.Unidade]; !ok {
					somaPortotalsalariomensalfolhaUnidadeFuncaoPdt[pdt.Total_mensal_folha][pdt.Unidade] = make(map[string]float64)
				}

			
				somaPortotalsalariomensalfolhaUnidadeFuncaoPdt[pdt.Total_mensal_folha][pdt.Unidade][pdt.Funcao]=0.0
				somaPortotalsalariomensalfolhaUnidadeFuncaoPdt[pdt.Total_mensal_folha][pdt.Unidade][pdt.Funcao] += pdt.Total_salario_folha + pdt.Encargosfo
			}
		
			
			for _, iadvh := range iadvhs {
               // soma=0.0
				if _, ok := somaPortotalsalariomensalfolhaUnidadeFuncaoPdt[iadvh.Total_mensal_folha]; !ok {
					somaPortotalsalariomensalfolhaUnidadeFuncaoPdt[iadvh.Total_mensal_folha] = make(map[string]map[string]float64)
				}
				if _, ok := somaPortotalsalariomensalfolhaUnidadeFuncaoPdt[iadvh.Total_mensal_folha][iadvh.Unidade]; !ok {
					somaPortotalsalariomensalfolhaUnidadeFuncaoPdt[iadvh.Total_mensal_folha][iadvh.Unidade] = make(map[string]float64)
				}

	
 
				
				somaPortotalsalariomensalfolhaUnidadeFuncaoPdt[iadvh.Total_mensal_folha][iadvh.Unidade][iadvh.Funcao] += iadvh.Encargosia + iadvh.Total_salario_folha
			}
		
		

            for i,pdt:= range pdts {
                
                pdts[i].Total_mensal_folha =somaPortotalsalariomensalfolhaUnidadeFuncaoPdt[pdt.Total_mensal_folha][pdt.Unidade][pdt.Funcao]

            }



            somaPortotalsalarioanualfolhaUnidadeFuncaoPdt := make(map[float64]map[string]map[string]float64)
		
		
for _, pdt := range pdts {
	if _, ok := somaPortotalsalarioanualfolhaUnidadeFuncaoPdt[pdt.Total_anual_folha]; !ok {
		somaPortotalsalarioanualfolhaUnidadeFuncaoPdt[pdt.Total_anual_folha] = make(map[string]map[string]float64)
	}
	if _, ok := somaPortotalsalarioanualfolhaUnidadeFuncaoPdt[pdt.Total_anual_folha][pdt.Unidade]; !ok {
		somaPortotalsalarioanualfolhaUnidadeFuncaoPdt[pdt.Total_anual_folha][pdt.Unidade] = make(map[string]float64)
	}
	somaPortotalsalarioanualfolhaUnidadeFuncaoPdt[pdt.Total_anual_folha][pdt.Unidade][pdt.Funcao]=0.0
	somaPortotalsalarioanualfolhaUnidadeFuncaoPdt[pdt.Total_anual_folha][pdt.Unidade][pdt.Funcao] += (pdt.Total_mensal_folha)*12
   // somaPortotalsalarioanualfolhaUnidadeFuncaoPdt[pdt.Total_anual_folha][pdt.Unidade][pdt.Funcao]
}


for _, iadvh := range iadvhs {
    
	if _, ok := somaPortotalsalarioanualfolhaUnidadeFuncaoPdt[iadvh.Total_anual_folha]; !ok {
		somaPortotalsalarioanualfolhaUnidadeFuncaoPdt[iadvh.Total_anual_folha] = make(map[string]map[string]float64)
	}
	if _, ok := somaPortotalsalarioanualfolhaUnidadeFuncaoPdt[iadvh.Total_anual_folha][iadvh.Unidade]; !ok {
		somaPortotalsalarioanualfolhaUnidadeFuncaoPdt[iadvh.Total_anual_folha][iadvh.Unidade] = make(map[string]float64)
	}

	somaPortotalsalarioanualfolhaUnidadeFuncaoPdt[iadvh.Total_anual_folha][iadvh.Unidade][iadvh.Funcao] += (iadvh.Total_mensal_folha)*12
    //soma1= somaPortotalsalarioanualfolhaUnidadeFuncaoPdt[iadvh.Total_anual_folha][iadvh.Unidade][iadvh.Funcao]
}



for i,pdt:= range pdts{
    pdts[i].Total_anual_folha=somaPortotalsalarioanualfolhaUnidadeFuncaoPdt[pdt.Total_anual_folha][pdt.Unidade][pdt.Funcao]

}

        
    
    config.DB.Save(&pdts)

}





func worker(wg *sync.WaitGroup, done chan bool, calculos ...func()) {
	defer wg.Done()

	// Chama todas as funções de cálculo
	for _, calculo := range calculos {
		calculo()
	}

	done <- true
}


// @Summary Returns the specific unit with information for the pdt
// @Description Returns the unit
// @Accept  json
// @Produce  json
// @Param unidade body string true "Unidade"
// @Success 200 {array} models.PDT
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /unidade [post]
func GetUserss(context *gin.Context) {
    var pdt []models.PDT
	var pdts []models.PDT
	config.DB.Find(&pdts)

	var iadvhs []models.Iadvh
	config.DB.Find(&iadvhs)
//////////////////////////////////////////////////////////////////////////////
   //post

   done := make(chan bool)
   var wg sync.WaitGroup
   numGoroutines := 16
   wg.Add(numGoroutines)

   // Inicia as goroutines
   for i := 0; i < numGoroutines; i++ {
       go worker(&wg, done,CalcAdicionalnoturno,CalcAuxilioialimentacao,CalcGratificacao,CalcHoraExtra,CalcSalariobase,CalcPordeducao,CalcInsalubridade,CalcPorencargos,CalcDsr)
   }

   // Aguarda a conclusão de todas as goroutines
   go func() {
       wg.Wait()
       close(done)
   }()

   // Aguarda o canal done ser fechado
   <-done


CalcTotalSalarioFolha()
Calculo()
   var requestData struct {
     Unidade string `json:"unidade" binding:"required"`
}

if err := context.ShouldBindJSON(&requestData); err != nil {
	context.JSON(http.StatusBadRequest, gin.H{"error": "Por favor, forneça a unidade no corpo da solicitação"})
	return
}



if err := config.DB.Where("Unidade = ?", requestData.Unidade).Find(&pdt).Error; err != nil {
	context.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao consultar o banco de dados"})
	return
}

if len(pdt) == 0 {
	context.JSON(http.StatusNotFound, gin.H{"error": "Nenhum registro encontrado para a unidade especificada"})
	return
}

// Usando um mapa para armazenar funções únicas
funcoesUnicas := make(map[string]struct{})

// Filtra funções únicas e atualiza o slice pdt
var pdtResultado []models.PDT
//var iadvhresultado []models.Iadvh


for _, item := range pdt {
	if _, ok := funcoesUnicas[item.Funcao]; !ok {
		funcoesUnicas[item.Funcao] = struct{}{}
		pdtResultado = append(pdtResultado, item)
	}
}

/*
if len(pdt) == 0{
    for _, item := range iadvhs {
        if _, ok := funcoesUnicas[item.Funcao]; !ok {
            funcoesUnicas[item.Funcao] = struct{}{}
            iadvhresultado = append(iadvhresultado, item)
        }
    }
}
*/
context.JSON(http.StatusOK, pdtResultado)
}