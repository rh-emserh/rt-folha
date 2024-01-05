package controller

import (
	"net/http" // Uncomment this line
	"github.com/gin-gonic/gin"
	"example.com/gogin/config"
	"example.com/gogin/models"


//"gorm.io/gorm"

//"fmt"
//"encoding/json"



)


func Calculo(c *gin.Context) {

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
			somaPoradicionalnoturnoUnidadeFuncaoPdt[pdt.Adicional_noturno_20][pdt.Unidade][pdt.Funcao] += pdt.Adicional_noturno_20+pdt.Adicional_noturno_30+pdt.Adicional_noturno_20_valor+pdt.Media_adicional_noturno_salario_maternidade+pdt.Adicional_noturno_30_valor+pdt.Adicional_noturno_30_mes_anterior+pdt.Diferenca_adicional_noturno+pdt.Adicional_noturno_20_retroativo
		}
	
		for _, iadvh := range iadvhs {
			if _, ok := somaPoradicionalnoturnoUnidadeFuncaoPdt[iadvh.Adicional_noturno_mes_anterior]; !ok {
				somaPoradicionalnoturnoUnidadeFuncaoPdt[iadvh.Adicional_noturno_mes_anterior] = make(map[string]map[string]float64)
			}
			if _, ok := somaPoradicionalnoturnoUnidadeFuncaoPdt[iadvh.Adicional_noturno_mes_anterior][iadvh.Unidade]; !ok {
				somaPoradicionalnoturnoUnidadeFuncaoPdt[iadvh.Adicional_noturno_mes_anterior][iadvh.Unidade] = make(map[string]float64)
			}
			somaPoradicionalnoturnoUnidadeFuncaoPdt[iadvh.Adicional_noturno_mes_anterior][iadvh.Unidade][iadvh.Funcao] += iadvh.Adicional_noturno+iadvh.Adicional_noturno_mes_anterior
		}
	
		
		for i, pdt := range pdts {
			soma := somaPoradicionalnoturnoUnidadeFuncaoPdt[pdt.Adicional_noturno_20][pdt.Unidade][pdt.Funcao]
			pdts[i].Adicional_noturnofo = soma
		}
		
		config.DB.Save(&pdts)
}


/////////////////////////////////////////////////////////
//função de adicional noturno




func GetUserss(context *gin.Context) {
    var pdt []models.PDT
	var pdts []models.PDT
	config.DB.Find(&pdts)

	var iadvhs []models.Iadvh
	config.DB.Find(&iadvhs)




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
	for _, item := range pdt {
		if _, ok := funcoesUnicas[item.Funcao]; !ok {
			funcoesUnicas[item.Funcao] = struct{}{}
			pdtResultado = append(pdtResultado, item)
		}
	}

	context.JSON(http.StatusOK, pdtResultado)

   /* // Analisa o JSON do corpo da solicitação e verifica se o campo "unidade" está presente
    if err := context.ShouldBindJSON(&requestData); err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": "Por favor, forneça a unidade no corpo da solicitação"})
        return
    }

    // Filtra os resultados com base na unidade
    if err := config.DB.Where("Unidade = ?", requestData.Unidade).Find(&pdt).Error; err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao consultar o banco de dados"})
        return
    }

    // Verifica se foram encontrados registros
    if len(pdt) == 0 {
        context.JSON(http.StatusNotFound, gin.H{"error": "Nenhum registro encontrado para a unidade especificada"})
        return
    }

    context.JSON(http.StatusOK, pdt)*/
}

/*
func Somaauxilioalimentacao(context *gin.Context) {
    var pdts []models.PDT
    config.DB.Find(&pdts)


    var iadvhs []models.Iadvh
    config.DB.Find(&iadvhs)


    somaPoradicionalnoturnoUnidadeFuncaoPdt := make(map[float64]map[string]map[string]float64)


    // Calcula a soma para pdts
    for _, pdt := range pdts {
        indice := pdt.Adicional_noturno_20 // Use um índice comum
        if _, ok := somaPoradicionalnoturnoUnidadeFuncaoPdt[indice]; !ok {
            somaPoradicionalnoturnoUnidadeFuncaoPdt[indice] = make(map[string]map[string]float64)
        }
        if _, ok := somaPoradicionalnoturnoUnidadeFuncaoPdt[indice][pdt.Unidade]; !ok {
            somaPoradicionalnoturnoUnidadeFuncaoPdt[indice][pdt.Unidade] = make(map[string]float64)
        }
        somaPoradicionalnoturnoUnidadeFuncaoPdt[indice][pdt.Unidade][pdt.Funcao] += pdt.Adicional_noturno_20 + pdt.Adicional_noturno_30 + pdt.Adicional_noturno_20_valor + pdt.Media_adicional_noturno_salario_maternidade + pdt.Adicional_noturno_30_valor + pdt.Adicional_noturno_30_mes_anterior + pdt.Diferenca_adicional_noturno + pdt.Adicional_noturno_20_retroativo
    }


    // Calcula a soma para iadvhs
    for _, iadvh := range iadvhs {
        indice := iadvh.Adicional_noturno_mes_anterior // Use um índice comum
        if _, ok := somaPoradicionalnoturnoUnidadeFuncaoPdt[indice]; !ok {
            somaPoradicionalnoturnoUnidadeFuncaoPdt[indice] = make(map[string]map[string]float64)
        }
        if _, ok := somaPoradicionalnoturnoUnidadeFuncaoPdt[indice][iadvh.Unidade]; !ok {
            somaPoradicionalnoturnoUnidadeFuncaoPdt[indice][iadvh.Unidade] = make(map[string]float64)
        }
        somaPoradicionalnoturnoUnidadeFuncaoPdt[indice][iadvh.Unidade][iadvh.Funcao] += iadvh.Adicional_noturno + iadvh.Adicional_noturno_mes_anterior
    }


    // Calcula a soma final e armazena em uma variável
 
//já foi



	somaPorauxilioalimentacaoUnidadeFuncaoPdt := make(map[float64]map[string]map[string]float64)


    // Calcula a soma para pdts
    for _, pdt := range pdts {
        indice := pdt.Adicional_noturno_20 // Use um índice comum
        if _, ok :=  somaPorauxilioalimentacaoUnidadeFuncaoPdt[indice]; !ok {
			somaPorauxilioalimentacaoUnidadeFuncaoPdt[indice] = make(map[string]map[string]float64)
        }
        if _, ok :=  somaPorauxilioalimentacaoUnidadeFuncaoPdt[indice][pdt.Unidade]; !ok {
			somaPorauxilioalimentacaoUnidadeFuncaoPdt[indice][pdt.Unidade] = make(map[string]float64)
        }
		somaPorauxilioalimentacaoUnidadeFuncaoPdt[indice][pdt.Unidade][pdt.Funcao] += pdt.Auxilio_alimentacao+pdt.Auxilio_alimentacao_retroativo
    }


    // Calcula a soma final e armazena em uma variável



//já foi




	somaPorpericulosidadeUnidadeFuncaoPdt  := make(map[float64]map[string]map[string]float64)


    // Calcula a soma para pdts
    for _, pdt := range pdts {
        indice := pdt.Adicional_noturno_20 // Use um índice comum
        if _, ok := somaPorpericulosidadeUnidadeFuncaoPdt [indice]; !ok {
            somaPorpericulosidadeUnidadeFuncaoPdt [indice] = make(map[string]map[string]float64)
        }
        if _, ok := somaPorpericulosidadeUnidadeFuncaoPdt [indice][pdt.Unidade]; !ok {
            somaPorpericulosidadeUnidadeFuncaoPdt [indice][pdt.Unidade] = make(map[string]float64)
        }
        somaPorpericulosidadeUnidadeFuncaoPdt [indice][pdt.Unidade][pdt.Funcao] += pdt.Periculosidade
    }


    // Calcula a soma para iadvhs
    for _, iadvh := range iadvhs {
        indice := iadvh.Adicional_noturno_mes_anterior // Use um índice comum
        if _, ok := somaPorpericulosidadeUnidadeFuncaoPdt [indice]; !ok {
            somaPorpericulosidadeUnidadeFuncaoPdt [indice] = make(map[string]map[string]float64)
        }
        if _, ok := somaPorpericulosidadeUnidadeFuncaoPdt [indice][iadvh.Unidade]; !ok {
            somaPorpericulosidadeUnidadeFuncaoPdt [indice][iadvh.Unidade] = make(map[string]float64)
        }
        somaPorpericulosidadeUnidadeFuncaoPdt [indice][iadvh.Unidade][iadvh.Funcao] += iadvh.Periculosidade
    }


    // Calcula a soma final e armazena em uma variável
  


	somaPorgratificacaoUnidadeFuncaoPdt   := make(map[float64]map[string]map[string]float64)


    // Calcula a soma para pdts
    for _, pdt := range pdts {
        indice := pdt.Adicional_noturno_20 // Use um índice comum
        if _, ok := somaPorgratificacaoUnidadeFuncaoPdt  [indice]; !ok {
            somaPorgratificacaoUnidadeFuncaoPdt  [indice] = make(map[string]map[string]float64)
        }
        if _, ok := somaPorgratificacaoUnidadeFuncaoPdt  [indice][pdt.Unidade]; !ok {
            somaPorgratificacaoUnidadeFuncaoPdt  [indice][pdt.Unidade] = make(map[string]float64)
        }
        somaPorgratificacaoUnidadeFuncaoPdt  [indice][pdt.Unidade][pdt.Funcao] += pdt.Gratificacao_funcao_retroativo+pdt.Gratificacao_funcao+pdt.Gratificacao_dedicacao_exclusiva+pdt.Gratificacao_produtividade+pdt.Gratificacao_funcao_supervisor_enfermagem+pdt.Adicional_graduacao_15_tecnico_enfermage+pdt.Gratificacao_5_sobre_salario_base+pdt.Prorrogacao_media_deducao_exclusiva_salario_maternidade+pdt.Gratificacao_funcao_instrumentacao_cirurgico+pdt.Gratificacao_por_interiorizacao+pdt.Gratificacao+pdt.Adicional_graduacao_15_tecnico_enfermagem_licenca_maternidade+pdt.Media_gratificacao_funcao_prorrogacao_maternidade+pdt.Gratificacao_funcao_supervisor_enfermagem_sl
    }


    // Calcula a soma para iadvhs
    for _, iadvh := range iadvhs {
        indice := iadvh.Adicional_noturno_mes_anterior // Use um índice comum
        if _, ok := somaPorgratificacaoUnidadeFuncaoPdt  [indice]; !ok {
            somaPorgratificacaoUnidadeFuncaoPdt  [indice] = make(map[string]map[string]float64)
        }
        if _, ok := somaPorgratificacaoUnidadeFuncaoPdt  [indice][iadvh.Unidade]; !ok {
            somaPorgratificacaoUnidadeFuncaoPdt  [indice][iadvh.Unidade] = make(map[string]float64)
        }
        somaPorgratificacaoUnidadeFuncaoPdt  [indice][iadvh.Unidade][iadvh.Funcao] += iadvh.Gratificacao+iadvh.Gratificacao_cargo_confianca+iadvh.Gratificacao_indenizatoria+iadvh.Gratificacao_funcao_instrumentacao_cirurgico+iadvh.Gratificacao_mes_anterior
    }


    // Calcula a soma final e armazena em uma variável




	
	somaPorhoraextraUnidadeFuncaoPdt    := make(map[float64]map[string]map[string]float64)


    // Calcula a soma para pdts
    for _, pdt := range pdts {
        indice := pdt.Adicional_noturno_20 // Use um índice comum
        if _, ok := somaPorhoraextraUnidadeFuncaoPdt   [indice]; !ok {
            somaPorhoraextraUnidadeFuncaoPdt   [indice] = make(map[string]map[string]float64)
        }
        if _, ok := somaPorhoraextraUnidadeFuncaoPdt   [indice][pdt.Unidade]; !ok {
            somaPorhoraextraUnidadeFuncaoPdt   [indice][pdt.Unidade] = make(map[string]float64)
        }
        somaPorhoraextraUnidadeFuncaoPdt   [indice][pdt.Unidade][pdt.Funcao] += pdt.Media_horas_extras_prorrogacao_salario_maternidade+pdt.Hora_extra_50+pdt.Hora_extra_sobre_adicionais_50+pdt.Hora_extra_sobre_adicionais_60+pdt.Hora_extra_sobre_adicionais_100+pdt.Dsr_sobre_horas_extras+pdt.Dsr_sem_horas_extras_retroativas+pdt.Hora_extra_100+pdt.Hora_extra_retroativa_60+pdt.Hora_extra_50_valor+pdt.Hora_extra_100_mes_anterior+pdt.Hora_extra_50_retroativo+pdt.Hora_extra_100_valor+pdt.Hora_extra_60_valor
    }


    // Calcula a soma para iadvhs
    for _, iadvh := range iadvhs {
        indice := iadvh.Adicional_noturno_mes_anterior // Use um índice comum
        if _, ok := somaPorhoraextraUnidadeFuncaoPdt  [indice]; !ok {
            somaPorhoraextraUnidadeFuncaoPdt   [indice] = make(map[string]map[string]float64)
        }
        if _, ok := somaPorhoraextraUnidadeFuncaoPdt   [indice][iadvh.Unidade]; !ok {
            somaPorhoraextraUnidadeFuncaoPdt   [indice][iadvh.Unidade] = make(map[string]float64)
        }
        somaPorhoraextraUnidadeFuncaoPdt   [indice][iadvh.Unidade][iadvh.Funcao] += iadvh.Hora_extra_100+iadvh.Hora_extra_50+iadvh.Hora_extra_60+iadvh.Hora_extra_60_mes_anterior+iadvh.Horas_extras_100_mes_anterior+iadvh.Horas_extras_50_mes_anterior
    }


    // Calcula a soma final e armazena em uma variável
    


	somaPorsalariobaseUnidadeFuncaoPdt    := make(map[float64]map[string]map[string]float64)


    // Calcula a soma para pdts
    for _, pdt := range pdts {
        indice := pdt.Adicional_noturno_20 // Use um índice comum
        if _, ok := somaPorsalariobaseUnidadeFuncaoPdt   [indice]; !ok {
            somaPorsalariobaseUnidadeFuncaoPdt   [indice] = make(map[string]map[string]float64)
        }
        if _, ok := somaPorsalariobaseUnidadeFuncaoPdt   [indice][pdt.Unidade]; !ok {
            somaPorsalariobaseUnidadeFuncaoPdt   [indice][pdt.Unidade] = make(map[string]float64)
        }
        somaPorsalariobaseUnidadeFuncaoPdt   [indice][pdt.Unidade][pdt.Funcao] += pdt.Salario_base+pdt.Jovem_aprendiz+pdt.Diferenca_salarial+pdt.Devolucao_falta+pdt.Reembolso_faltas+pdt.Salario_substituicao+pdt.Salario_substituicao_mes_anterior+pdt.Estorno_desconto_vale_transporte+pdt.Dias_atestado+pdt.Prorrogacao_licenca_maternidade+pdt.Licenca_remunerada+pdt.Reembolso_educacional+pdt.Insuficiencia_saldo_provento+pdt.Saldo_salario+pdt.Saldo_salario_horista
    }


    // Calcula a soma para iadvhs
    for _, iadvh := range iadvhs {
        indice := iadvh.Adicional_noturno_mes_anterior // Use um índice comum
        if _, ok := somaPorsalariobaseUnidadeFuncaoPdt    [indice]; !ok {
            somaPorsalariobaseUnidadeFuncaoPdt   [indice] = make(map[string]map[string]float64)
        }
        if _, ok := somaPorsalariobaseUnidadeFuncaoPdt   [indice][iadvh.Unidade]; !ok {
            somaPorsalariobaseUnidadeFuncaoPdt   [indice][iadvh.Unidade] = make(map[string]float64)
        }
        somaPorsalariobaseUnidadeFuncaoPdt   [indice][iadvh.Unidade][iadvh.Funcao] += iadvh.Atestado_medico+iadvh.Devolucao_desconto_indevido+iadvh.Devolucao_faltas+iadvh.Desconto_plano_saude_hapvida+iadvh.Desconto_plano_odontologico_integral+iadvh.Dias_trabalhados+iadvh.Diferenca_salarial+iadvh.Horas_noturnas_reduzidas+iadvh.Insuficiencia_saldo_proventos+iadvh.Licenca_paternidade+iadvh.Saldo_salario
    }


    // Calcula a soma final e armazena em uma variável
  


	somaPordeducaoUnidadeFuncaoPdt     := make(map[float64]map[string]map[string]float64)


    // Calcula a soma para pdts
    for _, pdt := range pdts {
        indice := pdt.Adicional_noturno_20 // Use um índice comum
        if _, ok := somaPordeducaoUnidadeFuncaoPdt    [indice]; !ok {
            somaPordeducaoUnidadeFuncaoPdt    [indice] = make(map[string]map[string]float64)
        }
        if _, ok := somaPordeducaoUnidadeFuncaoPdt    [indice][pdt.Unidade]; !ok {
            somaPordeducaoUnidadeFuncaoPdt    [indice][pdt.Unidade] = make(map[string]float64)
        }
        somaPordeducaoUnidadeFuncaoPdt    [indice][pdt.Unidade][pdt.Funcao] += pdt.Gratificacao_produtor_salario_maternidade+pdt.Gratificacao_dedicacao_exclusiva_salario_maternidade+pdt.Gratificacao_interiorizacao_salario_maternidade+pdt.Salario_maternidade_jovem_aprendiz+pdt.Salario_maternidade+pdt.Media_horas_extras_salario_maternidade+pdt.Media_insalubridade_salario_maternidade+pdt.Gratificacao_sobre_maternidade_supervisor_enfermagem+pdt.Adicional_graduacao_15_tecnico_enfermagem_licenca_maternidade+pdt.Media_gratificacao_sobre_maternidade+pdt.Salario_familia+pdt.Salario_familia_indenizado 
    }


    // Calcula a soma para iadvhs
    for _, iadvh := range iadvhs {
        indice := iadvh.Adicional_noturno_mes_anterior // Use um índice comum
        if _, ok := somaPordeducaoUnidadeFuncaoPdt    [indice]; !ok {
            somaPordeducaoUnidadeFuncaoPdt    [indice] = make(map[string]map[string]float64)
        }
        if _, ok := somaPordeducaoUnidadeFuncaoPdt    [indice][iadvh.Unidade]; !ok {
            somaPordeducaoUnidadeFuncaoPdt    [indice][iadvh.Unidade] = make(map[string]float64)
        }
        somaPordeducaoUnidadeFuncaoPdt    [indice][iadvh.Unidade][iadvh.Funcao] += iadvh.Salario_maternidade+iadvh.Salario_familia+iadvh.Media_salario_maternidade
    }


    // Calcula a soma final e armazena em uma variável
   


	somaPorencargosUnidadeFuncaoPdt      := make(map[float64]map[string]map[string]float64)


    // Calcula a soma para pdts
    for _, pdt := range pdts {
        indice := pdt.Adicional_noturno_20 // Use um índice comum
        if _, ok := somaPorencargosUnidadeFuncaoPdt     [indice]; !ok {
            somaPorencargosUnidadeFuncaoPdt     [indice] = make(map[string]map[string]float64)
        }
        if _, ok := somaPorencargosUnidadeFuncaoPdt     [indice][pdt.Unidade]; !ok {
            somaPorencargosUnidadeFuncaoPdt   [indice][pdt.Unidade] = make(map[string]float64)
        }
        somaPorencargosUnidadeFuncaoPdt   [indice][pdt.Unidade][pdt.Funcao] += (pdt.Fgts_folha+pdt.Ferias+pdt.Ferias_nao_gozadas+pdt.Indenizacao_ferias_nao_gozadas+pdt.Ferias_sobre_aviso_indenizado+pdt.Aviso_previo_indenizado+pdt.Decimo_terceiro_sobre_aviso_indenizado+pdt.Ferias_vencidas+pdt.Ferias_proporcionais+pdt.Fgts_recisao+pdt.Fgts_decimo_terceiro_salario_recisao+pdt.Um_terco_ferias_recisao+pdt.Um_terco_ferias+pdt.Fgts_ferias) - (pdt.Gratificacao_produtor_salario_maternidade+pdt.Gratificacao_dedicacao_exclusiva_salario_maternidade+pdt.Gratificacao_interiorizacao_salario_maternidade+pdt.Salario_maternidade_jovem_aprendiz+pdt.Salario_maternidade+pdt.Media_horas_extras_salario_maternidade+pdt.Media_insalubridade_salario_maternidade+pdt.Gratificacao_sobre_maternidade_supervisor_enfermagem+pdt.Adicional_graduacao_15_tecnico_enfermagem_licenca_maternidade+pdt.Media_gratificacao_sobre_maternidade+pdt.Salario_familia+pdt.Salario_familia_indenizado )
    }



    // Calcula a soma para iadvhs
    for _, iadvh := range iadvhs {
        indice := iadvh.Adicional_noturno_mes_anterior // Use um índice comum
        if _, ok := somaPorencargosUnidadeFuncaoPdt    [indice]; !ok {
            somaPorencargosUnidadeFuncaoPdt     [indice] = make(map[string]map[string]float64)
        }
        if _, ok := somaPorencargosUnidadeFuncaoPdt     [indice][iadvh.Unidade]; !ok {
            somaPorencargosUnidadeFuncaoPdt     [indice][iadvh.Unidade] = make(map[string]float64)
        }
        somaPorencargosUnidadeFuncaoPdt     [indice][iadvh.Unidade][iadvh.Funcao] += (iadvh.Um_terco_ferias_mes+iadvh.Um_terco_ferias_proximo_mes+iadvh.Decimo_terceiro_indenizado+iadvh.Decimo_terceiro_licenca_maternidade+iadvh.Decimo_terceiro_recisao+iadvh.Adicional_um_terco_ferias+iadvh.Adicional_um_terco_ferias_proporcional_recisao+iadvh.Aviso_previo_indenizado+iadvh.Ferias_indenizadas+iadvh.Ferias_mes+iadvh.Ferias_proximo_mes+iadvh.Ferias_proporcionais+iadvh.Media_decimo_terceiro_rescisao+iadvh.Media_aviso_previo+iadvh.Media_ferias+iadvh.Media_ferias_proporcionais+iadvh.Multa_art_479_clt)- ( iadvh.Salario_maternidade+iadvh.Salario_familia+iadvh.Media_salario_maternidade)
    }


    // Calcula a soma final e armazena em uma variável
    var resultadoCalculo []map[string]interface{}
    for _, pdt := range pdts {
        indice := pdt.Adicional_noturno_20
        soma := somaPorencargosUnidadeFuncaoPdt    [indice][pdt.Unidade][pdt.Funcao]
        resultadoCalculo = append(resultadoCalculo, map[string]interface{}{
            "unidade": pdt.Unidade,
            "funcao":  pdt.Funcao,
            "adicional_noturno":somaPoradicionalnoturnoUnidadeFuncaoPdt[indice][pdt.Unidade][pdt.Funcao]idadeFuncaoPdt,
			"auxilio_alimentacao": somaPorpericulosidadeUnidadeFuncaoPdt[indice][pdt.Unidade][pdt.Funcao],
			"periculosidade":somaPorpericulosidadeUnidadeFuncaoPdt[indice][pdt.Unidade][pdt.Funcao],
			"gratificacao":somaPorgratificacaoUnidadeFuncaoPdt[indice][pdt.Unidade][pdt.Funcao] ,
			"hora_extra":somaPorhoraextraUnidadeFuncaoPdt[indice][pdt.Unidade][pdt.Funcao], 
			"salario_folha": somaPorsalariobaseUnidadeFuncaoPdt[indice][pdt.Unidade][pdt.Funcao],
			"deducao_folha": somaPordeducaoUnidadeFuncaoPdt[indice][pdt.Unidade][pdt.Funcao], 
			"soma_encargos": soma, 
			

        })
    }


    // Retorna o resultado sem persistir no banco de dados
    context.JSON(200, gin.H{
        "message": "Resultado do cálculo de Adicional_noturnofo",
        "result":  resultadoCalculo,
    })
}
*/