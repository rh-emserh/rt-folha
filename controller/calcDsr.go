package controller

import (
	
	"example.com/gogin/config"
	"example.com/gogin/models"

//"gorm.io/gorm"

//"encoding/json"

//"sync"

)


func CalcDsr() {
    //var pdt []models.PDT
	var pdts []models.PDT
	config.DB.Find(&pdts)

	var iadvhs []models.Iadvh
	config.DB.Find(&iadvhs)



    somaPorDsrUnidadeFuncaoPdt := make(map[float64]map[string]map[string]float64)


	for _, pdt := range pdts {
        //soma2=0.0
		if _, ok := somaPorDsrUnidadeFuncaoPdt[pdt.Aviso_previo_indenizado]; !ok {
			somaPorDsrUnidadeFuncaoPdt[pdt.Aviso_previo_indenizado] = make(map[string]map[string]float64)
		}
		if _, ok := somaPorDsrUnidadeFuncaoPdt[pdt.Aviso_previo_indenizado][pdt.Unidade]; !ok {
			somaPorDsrUnidadeFuncaoPdt[pdt.Aviso_previo_indenizado][pdt.Unidade] = make(map[string]float64)
		}
		somaPorDsrUnidadeFuncaoPdt[pdt.Aviso_previo_indenizado][pdt.Unidade][pdt.Funcao] += pdt.Dsr_sobre_adicional_noturno + pdt.Dsr_sem_adicional_noturno_mes_anterior
	//somaPorDsrUnidadeFuncaoPdt[pdt.Aviso_previo_indenizado][pdt.Unidade][pdt.Funcao]
    }

	for _, iadvh := range iadvhs {
        
		if _, ok := somaPorDsrUnidadeFuncaoPdt[iadvh.Aviso_previo_indenizado]; !ok {
			somaPorDsrUnidadeFuncaoPdt[iadvh.Aviso_previo_indenizado] = make(map[string]map[string]float64)
		}
		if _, ok := somaPorDsrUnidadeFuncaoPdt[iadvh.Aviso_previo_indenizado][iadvh.Unidade]; !ok {
			somaPorDsrUnidadeFuncaoPdt[iadvh.Aviso_previo_indenizado][iadvh.Unidade] = make(map[string]float64)
		} 

		somaPorDsrUnidadeFuncaoPdt[iadvh.Aviso_previo_indenizado][iadvh.Unidade][iadvh.Funcao] += iadvh.Descanso_semanal_remunerado
	   //somaPorDsrUnidadeFuncaoPdt[iadvh.Aviso_previo_indenizado][iadvh.Unidade][iadvh.Funcao]
    }


	

    for i,pdt:= range pdts {
        pdts[i].Dsrfo = 	somaPorDsrUnidadeFuncaoPdt[pdt.Aviso_previo_indenizado][pdt.Unidade][pdt.Funcao]

    }

	config.DB.Save(&pdts)
}