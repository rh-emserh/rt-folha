package controller

import (
	"github.com/gin-gonic/gin"
	"example.com/gogin/config"
	"example.com/gogin/models"
	"net/http" 
//"gorm.io/gorm"


//"encoding/json"

//"sync"

)



func CalcAdicionalnoturno(context *gin.Context) {
    var pdt []models.PDT
	var pdts []models.PDT
	config.DB.Find(&pdts)

	var iadvhs []models.Iadvh
	config.DB.Find(&iadvhs)
	
	
	
	
		somaPoradicionalnoturnoUnidadeFuncaoPdt := make(map[float64]map[string]map[string]float64)

   
		for _, pdt := range pdts {   
			 
			if _, ok := somaPoradicionalnoturnoUnidadeFuncaoPdt[pdt.Adicional_noturno_20]; !ok {
				somaPoradicionalnoturnoUnidadeFuncaoPdt[pdt.Adicional_noturno_20] = make(map[string]map[string]float64)
			}
			if _, ok := somaPoradicionalnoturnoUnidadeFuncaoPdt[pdt.Adicional_noturno_20][pdt.Unidade]; !ok {
				somaPoradicionalnoturnoUnidadeFuncaoPdt[pdt.Adicional_noturno_20][pdt.Unidade] = make(map[string]float64)
			}

			//somaPoradicionalnoturnoUnidadeFuncaoPdt[pdt.Adicional_noturno_20][pdt.Unidade][pdt.Funcao]=0.0
	
			somaPoradicionalnoturnoUnidadeFuncaoPdt[pdt.Adicional_noturno_20][pdt.Unidade][pdt.Funcao] += pdt.Adicional_noturno_20+pdt.Adicional_noturno_30+pdt.Adicional_noturno_20_valor+pdt.Media_adicional_noturno_salario_maternidade+pdt.Adicional_noturno_30_valor+pdt.Adicional_noturno_30_mes_anterior+pdt.Diferenca_adicional_noturno+pdt.Adicional_noturno_20_retroativo
           
		}
	/*
		for _, iadvh := range iadvhs {

			if _, ok := somaPoradicionalnoturnoUnidadeFuncaoPdt[iadvh.Adicional_noturno_mes_anterior]; !ok {
				somaPoradicionalnoturnoUnidadeFuncaoPdt[iadvh.Adicional_noturno_mes_anterior] = make(map[string]map[string]float64)
			}
			if _, ok := somaPoradicionalnoturnoUnidadeFuncaoPdt[iadvh.Adicional_noturno_mes_anterior][iadvh.Unidade]; !ok {
				somaPoradicionalnoturnoUnidadeFuncaoPdt[iadvh.Adicional_noturno_mes_anterior][iadvh.Unidade] = make(map[string]float64)
			}	
		
			somaPoradicionalnoturnoUnidadeFuncaoPdt[iadvh.Adicional_noturno_mes_anterior][iadvh.Unidade][iadvh.Funcao] += iadvh.Adicional_noturno+iadvh.Adicional_noturno_mes_anterior
          
		}
	*/
            for i,pdt:= range pdts{
            
             pdts[i].Adicional_noturnofo=somaPoradicionalnoturnoUnidadeFuncaoPdt[pdt.Adicional_noturno_20][pdt.Unidade][pdt.Funcao]

            
   
			if (i+1)%330 == 0 {
				slice := pdts[i-229 : i+1]
				config.DB.Save(&slice)
			}
		}
		


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