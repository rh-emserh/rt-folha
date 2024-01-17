package controller

import (

	"example.com/gogin/config"
	"example.com/gogin/models"

//"gorm.io/gorm"

//"encoding/json"

//"sync"

)




func CalcPorencargos() {
    //var pdt []models.PDT
	var pdts []models.PDT
	config.DB.Find(&pdts)

	var iadvhs []models.Iadvh
	config.DB.Find(&iadvhs)



            somaPorencargosUnidadeFuncaoPdt := make(map[float64]map[string]map[string]float64)


	for _, pdt := range pdts {
       // soma2=0.0
		if _, ok := somaPorencargosUnidadeFuncaoPdt[pdt.Aviso_previo_indenizado]; !ok {
			somaPorencargosUnidadeFuncaoPdt[pdt.Aviso_previo_indenizado] = make(map[string]map[string]float64)
		}
		if _, ok := somaPorencargosUnidadeFuncaoPdt[pdt.Aviso_previo_indenizado][pdt.Unidade]; !ok {
			somaPorencargosUnidadeFuncaoPdt[pdt.Aviso_previo_indenizado][pdt.Unidade] = make(map[string]float64)
		}
		somaPorencargosUnidadeFuncaoPdt[pdt.Aviso_previo_indenizado][pdt.Unidade][pdt.Funcao] += (pdt.Fgts_folha+pdt.Ferias+pdt.Ferias_nao_gozadas+pdt.Indenizacao_ferias_nao_gozadas+pdt.Ferias_sobre_aviso_indenizado+pdt.Aviso_previo_indenizado+pdt.Decimo_terceiro_sobre_aviso_indenizado+pdt.Ferias_vencidas+pdt.Ferias_proporcionais+pdt.Fgts_recisao+pdt.Fgts_decimo_terceiro_salario_recisao+pdt.Um_terco_ferias_recisao+pdt.Um_terco_ferias+pdt.Fgts_ferias) - (pdt.Gratificacao_produtor_salario_maternidade+pdt.Gratificacao_dedicacao_exclusiva_salario_maternidade+pdt.Gratificacao_interiorizacao_salario_maternidade+pdt.Salario_maternidade_jovem_aprendiz+pdt.Salario_maternidade+pdt.Media_horas_extras_salario_maternidade+pdt.Media_insalubridade_salario_maternidade+pdt.Gratificacao_sobre_maternidade_supervisor_enfermagem+pdt.Adicional_graduacao_15_tecnico_enfermagem_licenca_maternidade+pdt.Media_gratificacao_sobre_maternidade+pdt.Salario_familia+pdt.Salario_familia_indenizado )
    }

	for _, iadvh := range iadvhs {
        //soma1=0.0
		if _, ok := somaPorencargosUnidadeFuncaoPdt[iadvh.Aviso_previo_indenizado]; !ok {
			somaPorencargosUnidadeFuncaoPdt[iadvh.Aviso_previo_indenizado] = make(map[string]map[string]float64)
		}
		if _, ok := somaPorencargosUnidadeFuncaoPdt[iadvh.Aviso_previo_indenizado][iadvh.Unidade]; !ok {
			somaPorencargosUnidadeFuncaoPdt[iadvh.Aviso_previo_indenizado][iadvh.Unidade] = make(map[string]float64)
		}
		somaPorencargosUnidadeFuncaoPdt[iadvh.Aviso_previo_indenizado][iadvh.Unidade][iadvh.Funcao] += (iadvh.Um_terco_ferias_mes+iadvh.Um_terco_ferias_proximo_mes+iadvh.Decimo_terceiro_indenizado+iadvh.Decimo_terceiro_licenca_maternidade+iadvh.Decimo_terceiro_recisao+iadvh.Adicional_um_terco_ferias+iadvh.Adicional_um_terco_ferias_proporcional_recisao+iadvh.Aviso_previo_indenizado+iadvh.Ferias_indenizadas+iadvh.Ferias_mes+iadvh.Ferias_proximo_mes+iadvh.Ferias_proporcionais+iadvh.Media_decimo_terceiro_rescisao+iadvh.Media_aviso_previo+iadvh.Media_ferias+iadvh.Media_ferias_proporcionais+iadvh.Multa_art_479_clt)- ( iadvh.Salario_maternidade+iadvh.Salario_familia+iadvh.Media_salario_maternidade)
	  //soma1=somaPorencargosUnidadeFuncaoPdt[iadvh.Aviso_previo_indenizado][iadvh.Unidade][iadvh.Funcao]
    }

    for i,pdt:= range pdts{

        pdts[i].Encargosfo =somaPorencargosUnidadeFuncaoPdt[pdt.Aviso_previo_indenizado][pdt.Unidade][pdt.Funcao]


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