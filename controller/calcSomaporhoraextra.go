package controller

import (
	
	"example.com/gogin/config"
	"example.com/gogin/models"

//"gorm.io/gorm"


//"encoding/json"

//"sync"

)


func CalcHoraExtra() {
    //var pdt []models.PDT
	var pdts []models.PDT
	config.DB.Find(&pdts)

	var iadvhs []models.Iadvh
	config.DB.Find(&iadvhs)

	somaPorhoraextraUnidadeFuncaoPdt := make(map[float64]map[string]map[string]float64)


	for _, pdt := range pdts {
		
		if _, ok := somaPorhoraextraUnidadeFuncaoPdt[pdt.Hora_extra_50]; !ok {
			somaPorhoraextraUnidadeFuncaoPdt[pdt.Hora_extra_50] = make(map[string]map[string]float64)
		}
		if _, ok := somaPorhoraextraUnidadeFuncaoPdt[pdt.Hora_extra_50][pdt.Unidade]; !ok {
			somaPorhoraextraUnidadeFuncaoPdt[pdt.Hora_extra_50][pdt.Unidade] = make(map[string]float64)
		}
		somaPorhoraextraUnidadeFuncaoPdt[pdt.Hora_extra_50][pdt.Unidade][pdt.Funcao] += pdt.Media_horas_extras_prorrogacao_salario_maternidade+pdt.Hora_extra_50+pdt.Hora_extra_sobre_adicionais_50+pdt.Hora_extra_sobre_adicionais_60+pdt.Hora_extra_sobre_adicionais_100+pdt.Dsr_sobre_horas_extras+pdt.Dsr_sem_horas_extras_retroativas+pdt.Hora_extra_100+pdt.Hora_extra_retroativa_60+pdt.Hora_extra_50_valor+pdt.Hora_extra_100_mes_anterior+pdt.Hora_extra_50_retroativo+pdt.Hora_extra_100_valor+pdt.Hora_extra_60_valor
       
	}

	for _, iadvh := range iadvhs {
		if _, ok := somaPorhoraextraUnidadeFuncaoPdt[iadvh.Hora_extra_50]; !ok {
			somaPorhoraextraUnidadeFuncaoPdt[iadvh.Hora_extra_50] = make(map[string]map[string]float64)
		}
		if _, ok := somaPorhoraextraUnidadeFuncaoPdt[iadvh.Hora_extra_50][iadvh.Unidade]; !ok {
			somaPorhoraextraUnidadeFuncaoPdt[iadvh.Hora_extra_50][iadvh.Unidade] = make(map[string]float64)
		}
		somaPorhoraextraUnidadeFuncaoPdt[iadvh.Hora_extra_50][iadvh.Unidade][iadvh.Funcao] += iadvh.Hora_extra_100+iadvh.Hora_extra_50+iadvh.Hora_extra_60+iadvh.Hora_extra_60_mes_anterior+iadvh.Horas_extras_100_mes_anterior+iadvh.Horas_extras_50_mes_anterior
       // fmt.Println("soma por hora extra iadvh %0.2f",somaPorhoraextraUnidadeFuncaoPdt[iadvh.Hora_extra_50][iadvh.Unidade][iadvh.Funcao])
          //soma2 = somaPorhoraextraUnidadeFuncaoPdt[iadvh.Hora_extra_50][iadvh.Unidade][iadvh.Funcao]
	}

	

    for i,pdt:= range pdts{
     
     pdts[i].Hora_extrafo =  somaPorhoraextraUnidadeFuncaoPdt[pdt.Hora_extra_50][pdt.Unidade][pdt.Funcao]
    	if (i+1)%3000 == 0 {
		slice := pdts[i-2999 : i+1]
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