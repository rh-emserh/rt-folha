package controller

import (
	
	"example.com/gogin/config"
	"example.com/gogin/models"

//"gorm.io/gorm"


//"encoding/json"

//"sync"

)




func CalcAuxilioialimentacao() {
    //var pdt []models.PDT
	var pdts []models.PDT
	config.DB.Find(&pdts)

	var iadvhs []models.Iadvh
	config.DB.Find(&iadvhs)
	
	

somaPorauxilioalimentacaoUnidadeFuncaoPdt := make(map[float64]map[string]map[string]float64)

for _, pdt := range pdts {
	if somaPorauxilioalimentacaoUnidadeFuncaoPdt[pdt.Auxilio_alimentacao_retroativo] == nil {
		somaPorauxilioalimentacaoUnidadeFuncaoPdt[pdt.Auxilio_alimentacao_retroativo] = make(map[string]map[string]float64)
	}
	if somaPorauxilioalimentacaoUnidadeFuncaoPdt[pdt.Auxilio_alimentacao_retroativo][pdt.Unidade] == nil {
		somaPorauxilioalimentacaoUnidadeFuncaoPdt[pdt.Auxilio_alimentacao_retroativo][pdt.Unidade] = make(map[string]float64)
	}

	somaPorauxilioalimentacaoUnidadeFuncaoPdt[pdt.Auxilio_alimentacao_retroativo][pdt.Unidade][pdt.Funcao] = 0.0

	somaPorauxilioalimentacaoUnidadeFuncaoPdt[pdt.Auxilio_alimentacao_retroativo][pdt.Unidade][pdt.Funcao] += pdt.Auxilio_alimentacao + pdt.Auxilio_alimentacao_retroativo
}

	



    for i, pdt := range pdts {
        soma := somaPorauxilioalimentacaoUnidadeFuncaoPdt[pdt.Auxilio_alimentacao_retroativo][pdt.Unidade][pdt.Funcao]
        pdts[i].Auxilio_alimentacaofo = soma
	}


	config.DB.Save(&pdts)
}