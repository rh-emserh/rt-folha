package controller

import (

	"example.com/gogin/config"
	"example.com/gogin/models"

//"gorm.io/gorm"


//"encoding/json"

//"sync"

)


func CalcGratificacao() {
    //var pdt []models.PDT
	var pdts []models.PDT
	config.DB.Find(&pdts)

	var iadvhs []models.Iadvh
	config.DB.Find(&iadvhs)
	


	somaPorgratificacaoUnidadeFuncaoPdt := make(map[float64]map[string]map[string]float64)


	for _, pdt := range pdts {
		
		if _, ok := somaPorgratificacaoUnidadeFuncaoPdt[pdt.Gratificacao_funcao_instrumentacao_cirurgico]; !ok {

			somaPorgratificacaoUnidadeFuncaoPdt[pdt.Gratificacao_funcao_instrumentacao_cirurgico] = make(map[string]map[string]float64)
		}
		if _, ok := somaPorgratificacaoUnidadeFuncaoPdt[pdt.Gratificacao_funcao_instrumentacao_cirurgico][pdt.Unidade]; !ok {
			somaPorgratificacaoUnidadeFuncaoPdt[pdt.Gratificacao_funcao_instrumentacao_cirurgico][pdt.Unidade] = make(map[string]float64)
		}
		somaPorgratificacaoUnidadeFuncaoPdt[pdt.Gratificacao_funcao_instrumentacao_cirurgico][pdt.Unidade][pdt.Funcao]=0.0
		somaPorgratificacaoUnidadeFuncaoPdt[pdt.Gratificacao_funcao_instrumentacao_cirurgico][pdt.Unidade][pdt.Funcao] += pdt.Gratificacao_funcao_retroativo+pdt.Gratificacao_funcao+pdt.Gratificacao_dedicacao_exclusiva+pdt.Gratificacao_produtividade+pdt.Gratificacao_funcao_supervisor_enfermagem+pdt.Adicional_graduacao_15_tecnico_enfermage+pdt.Gratificacao_5_sobre_salario_base+pdt.Prorrogacao_media_deducao_exclusiva_salario_maternidade+pdt.Gratificacao_funcao_instrumentacao_cirurgico+pdt.Gratificacao_por_interiorizacao+pdt.Gratificacao+pdt.Adicional_graduacao_15_tecnico_enfermagem_licenca_maternidade+pdt.Media_gratificacao_funcao_prorrogacao_maternidade+pdt.Gratificacao_funcao_supervisor_enfermagem_sl+pdt.Adicional_graduacao_15_tecnico_enfermage

	}

	for _, iadvh := range iadvhs {
		
		if _, ok := somaPorgratificacaoUnidadeFuncaoPdt[iadvh.Gratificacao_funcao_instrumentacao_cirurgico]; !ok {
			somaPorgratificacaoUnidadeFuncaoPdt[iadvh.Gratificacao_funcao_instrumentacao_cirurgico] = make(map[string]map[string]float64)
		}
		if _, ok := somaPorgratificacaoUnidadeFuncaoPdt[iadvh.Gratificacao_funcao_instrumentacao_cirurgico][iadvh.Unidade]; !ok {
			somaPorgratificacaoUnidadeFuncaoPdt[iadvh.Gratificacao_funcao_instrumentacao_cirurgico][iadvh.Unidade] = make(map[string]float64)
		}
		
		somaPorgratificacaoUnidadeFuncaoPdt[iadvh.Gratificacao_funcao_instrumentacao_cirurgico][iadvh.Unidade][iadvh.Funcao] += iadvh.Gratificacao+iadvh.Gratificacao_cargo_confianca+iadvh.Gratificacao_indenizatoria+iadvh.Gratificacao_funcao_instrumentacao_cirurgico+iadvh.Gratificacao_mes_anterior
        
	}

	
    for i,pdt:= range pdts{
        
     pdts[i].Gratificacaofo =somaPorgratificacaoUnidadeFuncaoPdt[pdt.Gratificacao_funcao_instrumentacao_cirurgico][pdt.Unidade][pdt.Funcao]
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

