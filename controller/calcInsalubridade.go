package controller

import (
	
	"example.com/gogin/config"
	"example.com/gogin/models"

//"gorm.io/gorm"


//"encoding/json"

//"sync"

)



func CalcInsalubridade() {
    //var pdt []models.PDT
	var pdts []models.PDT
	config.DB.Find(&pdts)

	var iadvhs []models.Iadvh
	config.DB.Find(&iadvhs)



    somaPorInsalubridadeUnidadeFuncaoPdt := make(map[float64]map[string]map[string]float64)
		
		
			for _, pdt := range pdts {
				if _, ok := somaPorInsalubridadeUnidadeFuncaoPdt[pdt.Insalubridade_40]; !ok {
					somaPorInsalubridadeUnidadeFuncaoPdt[pdt.Insalubridade_40] = make(map[string]map[string]float64)
				}
				if _, ok := somaPorInsalubridadeUnidadeFuncaoPdt[pdt.Insalubridade_40][pdt.Unidade]; !ok {
					somaPorInsalubridadeUnidadeFuncaoPdt[pdt.Insalubridade_40][pdt.Unidade] = make(map[string]float64)
				}
				somaPorInsalubridadeUnidadeFuncaoPdt[pdt.Insalubridade_40][pdt.Unidade][pdt.Funcao] += pdt.Insalubridade_40 + pdt.Insalubridade_sem_base_40+pdt.Insalubridade_sem_minimo_20+ pdt.Insalubridade_sem_minimo_20+pdt.Insalubridade_retroativa+pdt.Media_insalubridade_prorrogacao_salario_maternidade
	//	somaPorInsalubridadeUnidadeFuncaoPdt[pdt.Insalubridade_40][pdt.Unidade][pdt.Funcao]
            }
		
			
			for _, iadvh := range iadvhs {
               // soma1=0.0
				if _, ok := somaPorInsalubridadeUnidadeFuncaoPdt[iadvh.Insalubridade_40]; !ok {
					somaPorInsalubridadeUnidadeFuncaoPdt[iadvh.Insalubridade_40] = make(map[string]map[string]float64)
				}
				if _, ok := somaPorInsalubridadeUnidadeFuncaoPdt[iadvh.Insalubridade_40][iadvh.Unidade]; !ok {
					somaPorInsalubridadeUnidadeFuncaoPdt[iadvh.Insalubridade_40][iadvh.Unidade] = make(map[string]float64)
				}
				somaPorInsalubridadeUnidadeFuncaoPdt[iadvh.Insalubridade_40][iadvh.Unidade][iadvh.Funcao] += iadvh.Diferenca_insalubridade+ iadvh.Insalubridade_20+ iadvh.Insalubridade_40+iadvh.Insalubridade_mes_anterior_40
			   //soma1=somaPorInsalubridadeUnidadeFuncaoPdt[iadvh.Insalubridade_40][iadvh.Unidade][iadvh.Funcao]
            }
		
		

            for i,pdt:= range pdts {
                
                pdts[i].Insalubridade =somaPorInsalubridadeUnidadeFuncaoPdt[pdt.Insalubridade_40][pdt.Unidade][pdt.Funcao]

            	if (i+1)%5000 == 0 {
				slice := pdts[i-4999 : i+1]
				config.DB.Save(&slice)
				//fmt.Printf("Atualizadas %d linhas.\n", i+1)
			}
		}
	
		// Certifique-se de realizar uma última atualização em lote se necessário
		if len(pdts)%3000 != 0 {
			slice := pdts[len(pdts)-(len(pdts)%3000):]
			config.DB.Save(&slice)
			//fmt.Printf("Atualizadas %d linhas.\n", len(pdts)%300)
		}
			
			//config.DB.Save(&pdts)
}