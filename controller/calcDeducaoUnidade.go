package controller

import (

	"example.com/gogin/config"
	"example.com/gogin/models"

//"gorm.io/gorm"

//"encoding/json"

//"sync"

)

func CalcPordeducao() {
    //var pdt []models.PDT
	var pdts []models.PDT
	config.DB.Find(&pdts)

	var iadvhs []models.Iadvh
	config.DB.Find(&iadvhs)


	somaPordeducaoUnidadeFuncaoPdt := make(map[float64]map[string]map[string]float64)

	for _, pdt := range pdts {
		if _, ok := somaPordeducaoUnidadeFuncaoPdt[pdt.Salario_maternidade]; !ok {
			somaPordeducaoUnidadeFuncaoPdt[pdt.Salario_maternidade] = make(map[string]map[string]float64)
		}
		if _, ok := somaPordeducaoUnidadeFuncaoPdt[pdt.Salario_maternidade][pdt.Unidade]; !ok {
			somaPordeducaoUnidadeFuncaoPdt[pdt.Salario_maternidade][pdt.Unidade] = make(map[string]float64)
		}
		somaPordeducaoUnidadeFuncaoPdt[pdt.Salario_maternidade][pdt.Unidade][pdt.Funcao] += pdt.Gratificacao_produtor_salario_maternidade+pdt.Gratificacao_dedicacao_exclusiva_salario_maternidade+pdt.Gratificacao_interiorizacao_salario_maternidade+pdt.Salario_maternidade_jovem_aprendiz+pdt.Salario_maternidade+pdt.Media_horas_extras_salario_maternidade+pdt.Media_insalubridade_salario_maternidade+pdt.Gratificacao_sobre_maternidade_supervisor_enfermagem+pdt.Adicional_graduacao_15_tecnico_enfermagem_licenca_maternidade+pdt.Media_gratificacao_sobre_maternidade+pdt.Salario_familia+pdt.Salario_familia_indenizado  //aqui tem 12
	    //somaPordeducaoUnidadeFuncaoPdt[pdt.Salario_maternidade][pdt.Unidade][pdt.Funcao]
    }


	for _, iadvh := range iadvhs {
		if _, ok := somaPordeducaoUnidadeFuncaoPdt[iadvh.Salario_maternidade]; !ok {
			somaPordeducaoUnidadeFuncaoPdt[iadvh.Salario_maternidade] = make(map[string]map[string]float64)
		}
		if _, ok := somaPordeducaoUnidadeFuncaoPdt[iadvh.Salario_maternidade][iadvh.Unidade]; !ok {
			somaPordeducaoUnidadeFuncaoPdt[iadvh.Salario_maternidade][iadvh.Unidade] = make(map[string]float64)
		}
		somaPordeducaoUnidadeFuncaoPdt[iadvh.Salario_maternidade][iadvh.Unidade][iadvh.Funcao] += iadvh.Salario_maternidade+iadvh.Salario_familia+iadvh.Media_salario_maternidade
      //  soma1=somaPordeducaoUnidadeFuncaoPdt[iadvh.Salario_maternidade][iadvh.Unidade][iadvh.Funcao]
	}

	


    for i,pdt:= range pdts {
        
        pdts[i].Deducao_inssfo =somaPordeducaoUnidadeFuncaoPdt[pdt.Salario_maternidade][pdt.Unidade][pdt.Funcao]

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