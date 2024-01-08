package controller

import (
	"net/http" // Uncomment this line
	"github.com/gin-gonic/gin"
	"example.com/gogin/config"
	"example.com/gogin/models"

//"gorm.io/gorm"

"fmt"
//"encoding/json"

"sync"

)

var soma1 float64 
var soma2 float64
var soma3 float64


func Calculo() {
    //var pdt []models.PDT
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
            soma2=somaPoradicionalnoturnoUnidadeFuncaoPdt[pdt.Adicional_noturno_20][pdt.Unidade][pdt.Funcao]
		}
	
		for _, iadvh := range iadvhs {
			if _, ok := somaPoradicionalnoturnoUnidadeFuncaoPdt[iadvh.Adicional_noturno_mes_anterior]; !ok {
				somaPoradicionalnoturnoUnidadeFuncaoPdt[iadvh.Adicional_noturno_mes_anterior] = make(map[string]map[string]float64)
			}
			if _, ok := somaPoradicionalnoturnoUnidadeFuncaoPdt[iadvh.Adicional_noturno_mes_anterior][iadvh.Unidade]; !ok {
				somaPoradicionalnoturnoUnidadeFuncaoPdt[iadvh.Adicional_noturno_mes_anterior][iadvh.Unidade] = make(map[string]float64)
			}
			somaPoradicionalnoturnoUnidadeFuncaoPdt[iadvh.Adicional_noturno_mes_anterior][iadvh.Unidade][iadvh.Funcao] += iadvh.Adicional_noturno+iadvh.Adicional_noturno_mes_anterior
            soma1 =somaPoradicionalnoturnoUnidadeFuncaoPdt[iadvh.Adicional_noturno_mes_anterior][iadvh.Unidade][iadvh.Funcao]
		}
	
            for i:= range pdts{
                soma3= soma1 +soma2
             pdts[i].Adicional_noturnofo = soma3
            }




        


/////////////////////////////////////////////////////////
//função de adicional noturno







	somaPorauxilioalimentacaoUnidadeFuncaoPdt := make(map[float64]map[string]map[string]float64)


	for _, pdt := range pdts {
		if _, ok := somaPorauxilioalimentacaoUnidadeFuncaoPdt[pdt.Auxilio_alimentacao_retroativo]; !ok {
			somaPorauxilioalimentacaoUnidadeFuncaoPdt[pdt.Auxilio_alimentacao_retroativo] = make(map[string]map[string]float64)
		}
		if _, ok := somaPorauxilioalimentacaoUnidadeFuncaoPdt[pdt.Auxilio_alimentacao_retroativo][pdt.Unidade]; !ok {
			somaPorauxilioalimentacaoUnidadeFuncaoPdt[pdt.Auxilio_alimentacao_retroativo][pdt.Unidade] = make(map[string]float64)
		}
		somaPorauxilioalimentacaoUnidadeFuncaoPdt[pdt.Auxilio_alimentacao_retroativo][pdt.Unidade][pdt.Funcao] += pdt.Auxilio_alimentacao+pdt.Auxilio_alimentacao_retroativo
	}

	



    for i, pdt := range pdts {
        soma := somaPorauxilioalimentacaoUnidadeFuncaoPdt[pdt.Auxilio_alimentacao_retroativo][pdt.Unidade][pdt.Funcao]
        pdts[i].Auxilio_alimentacaofo = soma
    }
    
  


	somaPorpericulosidadeUnidadeFuncaoPdt := make(map[float64]map[string]map[string]float64)


	for _, pdt := range pdts {
		if _, ok := somaPorpericulosidadeUnidadeFuncaoPdt[pdt.Periculosidade]; !ok {
			somaPorpericulosidadeUnidadeFuncaoPdt[pdt.Periculosidade] = make(map[string]map[string]float64)
		}
		if _, ok := somaPorpericulosidadeUnidadeFuncaoPdt[pdt.Periculosidade][pdt.Unidade]; !ok {
			somaPorpericulosidadeUnidadeFuncaoPdt[pdt.Periculosidade][pdt.Unidade] = make(map[string]float64)
		}
		somaPorpericulosidadeUnidadeFuncaoPdt[pdt.Periculosidade][pdt.Unidade][pdt.Funcao] += pdt.Periculosidade
        soma2= somaPorpericulosidadeUnidadeFuncaoPdt[pdt.Periculosidade][pdt.Unidade][pdt.Funcao]
	}

	for _, iadvh := range iadvhs {
		if _, ok := somaPorpericulosidadeUnidadeFuncaoPdt[iadvh.Periculosidade]; !ok {
			somaPorpericulosidadeUnidadeFuncaoPdt[iadvh.Periculosidade] = make(map[string]map[string]float64)
		}
		if _, ok := somaPorpericulosidadeUnidadeFuncaoPdt[iadvh.Periculosidade][iadvh.Unidade]; !ok {
			somaPorpericulosidadeUnidadeFuncaoPdt[iadvh.Periculosidade][iadvh.Unidade] = make(map[string]float64)
		}
		somaPorpericulosidadeUnidadeFuncaoPdt[iadvh.Periculosidade][iadvh.Unidade][iadvh.Funcao] += iadvh.Periculosidade
        soma1= somaPorpericulosidadeUnidadeFuncaoPdt[iadvh.Periculosidade][iadvh.Unidade][iadvh.Funcao]
	}

    for i:= range pdts{
        soma3= soma1 +soma2
     pdts[i].Periculosidadefo = soma3
    }

    






///////////////////////////////////////////////////////////////////////////
//calculo gratificação



	

	somaPorgratificacaoUnidadeFuncaoPdt := make(map[float64]map[string]map[string]float64)


	for _, pdt := range pdts {
		if _, ok := somaPorgratificacaoUnidadeFuncaoPdt[pdt.Gratificacao_funcao_instrumentacao_cirurgico]; !ok {
			somaPorgratificacaoUnidadeFuncaoPdt[pdt.Gratificacao_funcao_instrumentacao_cirurgico] = make(map[string]map[string]float64)
		}
		if _, ok := somaPorgratificacaoUnidadeFuncaoPdt[pdt.Gratificacao_funcao_instrumentacao_cirurgico][pdt.Unidade]; !ok {
			somaPorgratificacaoUnidadeFuncaoPdt[pdt.Gratificacao_funcao_instrumentacao_cirurgico][pdt.Unidade] = make(map[string]float64)
		}
		somaPorgratificacaoUnidadeFuncaoPdt[pdt.Gratificacao_funcao_instrumentacao_cirurgico][pdt.Unidade][pdt.Funcao] += pdt.Gratificacao_funcao_retroativo+pdt.Gratificacao_funcao+pdt.Gratificacao_dedicacao_exclusiva+pdt.Gratificacao_produtividade+pdt.Gratificacao_funcao_supervisor_enfermagem+pdt.Adicional_graduacao_15_tecnico_enfermage+pdt.Gratificacao_5_sobre_salario_base+pdt.Prorrogacao_media_deducao_exclusiva_salario_maternidade+pdt.Gratificacao_funcao_instrumentacao_cirurgico+pdt.Gratificacao_por_interiorizacao+pdt.Gratificacao+pdt.Adicional_graduacao_15_tecnico_enfermagem_licenca_maternidade+pdt.Media_gratificacao_funcao_prorrogacao_maternidade+pdt.Gratificacao_funcao_supervisor_enfermagem_sl
        soma2 =somaPorgratificacaoUnidadeFuncaoPdt[pdt.Gratificacao_funcao_instrumentacao_cirurgico][pdt.Unidade][pdt.Funcao]
	}

	for _, iadvh := range iadvhs {
		if _, ok := somaPorgratificacaoUnidadeFuncaoPdt[iadvh.Gratificacao_funcao_instrumentacao_cirurgico]; !ok {
			somaPorgratificacaoUnidadeFuncaoPdt[iadvh.Gratificacao_funcao_instrumentacao_cirurgico] = make(map[string]map[string]float64)
		}
		if _, ok := somaPorgratificacaoUnidadeFuncaoPdt[iadvh.Gratificacao_funcao_instrumentacao_cirurgico][iadvh.Unidade]; !ok {
			somaPorgratificacaoUnidadeFuncaoPdt[iadvh.Gratificacao_funcao_instrumentacao_cirurgico][iadvh.Unidade] = make(map[string]float64)
		}
		somaPorgratificacaoUnidadeFuncaoPdt[iadvh.Gratificacao_funcao_instrumentacao_cirurgico][iadvh.Unidade][iadvh.Funcao] += iadvh.Gratificacao+iadvh.Gratificacao_cargo_confianca+iadvh.Gratificacao_indenizatoria+iadvh.Gratificacao_funcao_instrumentacao_cirurgico+iadvh.Gratificacao_mes_anterior
        soma1 =somaPorgratificacaoUnidadeFuncaoPdt[iadvh.Gratificacao_funcao_instrumentacao_cirurgico][iadvh.Unidade][iadvh.Funcao]
	}

	
    for i:= range pdts{
        soma3= soma1 +soma2
     pdts[i].Gratificacaofo = soma3
    }




////////////////////////////////////////////////////////////////////////////



	

	somaPorhoraextraUnidadeFuncaoPdt := make(map[float64]map[string]map[string]float64)


	for _, pdt := range pdts {
		if _, ok := somaPorhoraextraUnidadeFuncaoPdt[pdt.Hora_extra_50]; !ok {
			somaPorhoraextraUnidadeFuncaoPdt[pdt.Hora_extra_50] = make(map[string]map[string]float64)
		}
		if _, ok := somaPorhoraextraUnidadeFuncaoPdt[pdt.Hora_extra_50][pdt.Unidade]; !ok {
			somaPorhoraextraUnidadeFuncaoPdt[pdt.Hora_extra_50][pdt.Unidade] = make(map[string]float64)
		}
		somaPorhoraextraUnidadeFuncaoPdt[pdt.Hora_extra_50][pdt.Unidade][pdt.Funcao] += pdt.Media_horas_extras_prorrogacao_salario_maternidade+pdt.Hora_extra_50+pdt.Hora_extra_sobre_adicionais_50+pdt.Hora_extra_sobre_adicionais_60+pdt.Hora_extra_sobre_adicionais_100+pdt.Dsr_sobre_horas_extras+pdt.Dsr_sem_horas_extras_retroativas+pdt.Hora_extra_100+pdt.Hora_extra_retroativa_60+pdt.Hora_extra_50_valor+pdt.Hora_extra_100_mes_anterior+pdt.Hora_extra_50_retroativo+pdt.Hora_extra_100_valor+pdt.Hora_extra_60_valor
        soma2 =somaPorhoraextraUnidadeFuncaoPdt[pdt.Hora_extra_50][pdt.Unidade][pdt.Funcao]
	}

	for _, iadvh := range iadvhs {
		if _, ok := somaPorhoraextraUnidadeFuncaoPdt[iadvh.Hora_extra_50]; !ok {
			somaPorhoraextraUnidadeFuncaoPdt[iadvh.Hora_extra_50] = make(map[string]map[string]float64)
		}
		if _, ok := somaPorhoraextraUnidadeFuncaoPdt[iadvh.Hora_extra_50][iadvh.Unidade]; !ok {
			somaPorhoraextraUnidadeFuncaoPdt[iadvh.Hora_extra_50][iadvh.Unidade] = make(map[string]float64)
		}
		somaPorhoraextraUnidadeFuncaoPdt[iadvh.Hora_extra_50][iadvh.Unidade][iadvh.Funcao] += iadvh.Hora_extra_100+iadvh.Hora_extra_50+iadvh.Hora_extra_60+iadvh.Hora_extra_60_mes_anterior+iadvh.Horas_extras_100_mes_anterior+iadvh.Horas_extras_50_mes_anterior
        fmt.Println("soma por hora extra iadvh %0.2f",somaPorhoraextraUnidadeFuncaoPdt[iadvh.Hora_extra_50][iadvh.Unidade][iadvh.Funcao])
          soma2 = somaPorhoraextraUnidadeFuncaoPdt[iadvh.Hora_extra_50][iadvh.Unidade][iadvh.Funcao]
	}

	

    for i:= range pdts{
        soma3= soma1 +soma2
     pdts[i].Hora_extrafo = soma3
    }
    

    


///////////////////////////////////////////////////////////////////////////
//calculo salario base




	somaPorsalariobaseUnidadeFuncaoPdt := make(map[float64]map[string]map[string]float64)

	for _, pdt := range pdts {
		if _, ok := somaPorsalariobaseUnidadeFuncaoPdt[pdt.Saldo_salario]; !ok {
			somaPorsalariobaseUnidadeFuncaoPdt[pdt.Saldo_salario] = make(map[string]map[string]float64)
		}
		if _, ok := somaPorsalariobaseUnidadeFuncaoPdt[pdt.Saldo_salario][pdt.Unidade]; !ok {
			somaPorsalariobaseUnidadeFuncaoPdt[pdt.Saldo_salario][pdt.Unidade] = make(map[string]float64)
		}
		somaPorsalariobaseUnidadeFuncaoPdt[pdt.Saldo_salario][pdt.Unidade][pdt.Funcao] += pdt.Salario_base+pdt.Jovem_aprendiz+pdt.Diferenca_salarial+pdt.Devolucao_falta+pdt.Reembolso_faltas+pdt.Salario_substituicao+pdt.Salario_substituicao_mes_anterior+pdt.Estorno_desconto_vale_transporte+pdt.Dias_atestado+pdt.Prorrogacao_licenca_maternidade+pdt.Licenca_remunerada+pdt.Reembolso_educacional+pdt.Insuficiencia_saldo_provento+pdt.Saldo_salario+pdt.Saldo_salario_horista
        soma2 = somaPorsalariobaseUnidadeFuncaoPdt[pdt.Saldo_salario][pdt.Unidade][pdt.Funcao]
	}


	for _, iadvh := range iadvhs {
		if _, ok := somaPorsalariobaseUnidadeFuncaoPdt[iadvh.Saldo_salario]; !ok {
			somaPorsalariobaseUnidadeFuncaoPdt[iadvh.Saldo_salario] = make(map[string]map[string]float64)
		}
		if _, ok := somaPorsalariobaseUnidadeFuncaoPdt[iadvh.Saldo_salario][iadvh.Unidade]; !ok {
			somaPorsalariobaseUnidadeFuncaoPdt[iadvh.Saldo_salario][iadvh.Unidade] = make(map[string]float64)
		}
		somaPorsalariobaseUnidadeFuncaoPdt[iadvh.Saldo_salario][iadvh.Unidade][iadvh.Funcao] += iadvh.Atestado_medico+iadvh.Devolucao_desconto_indevido+iadvh.Devolucao_faltas+iadvh.Desconto_plano_saude_hapvida+iadvh.Desconto_plano_odontologico_integral+iadvh.Dias_trabalhados+iadvh.Diferenca_salarial+iadvh.Horas_noturnas_reduzidas+iadvh.Insuficiencia_saldo_proventos+iadvh.Licenca_paternidade+iadvh.Saldo_salario
        soma1 = somaPorsalariobaseUnidadeFuncaoPdt[iadvh.Saldo_salario][iadvh.Unidade][iadvh.Funcao]
	}





    
    for i:= range pdts {
        pdts[i].Salariobasefo = soma1 +soma2
    }
    
 

	








	somaPordeducaoUnidadeFuncaoPdt := make(map[float64]map[string]map[string]float64)

	for _, pdt := range pdts {
		if _, ok := somaPordeducaoUnidadeFuncaoPdt[pdt.Salario_maternidade]; !ok {
			somaPordeducaoUnidadeFuncaoPdt[pdt.Salario_maternidade] = make(map[string]map[string]float64)
		}
		if _, ok := somaPordeducaoUnidadeFuncaoPdt[pdt.Salario_maternidade][pdt.Unidade]; !ok {
			somaPordeducaoUnidadeFuncaoPdt[pdt.Salario_maternidade][pdt.Unidade] = make(map[string]float64)
		}
		somaPordeducaoUnidadeFuncaoPdt[pdt.Salario_maternidade][pdt.Unidade][pdt.Funcao] += pdt.Gratificacao_produtor_salario_maternidade+pdt.Gratificacao_dedicacao_exclusiva_salario_maternidade+pdt.Gratificacao_interiorizacao_salario_maternidade+pdt.Salario_maternidade_jovem_aprendiz+pdt.Salario_maternidade+pdt.Media_horas_extras_salario_maternidade+pdt.Media_insalubridade_salario_maternidade+pdt.Gratificacao_sobre_maternidade_supervisor_enfermagem+pdt.Adicional_graduacao_15_tecnico_enfermagem_licenca_maternidade+pdt.Media_gratificacao_sobre_maternidade+pdt.Salario_familia+pdt.Salario_familia_indenizado  //aqui tem 12
	     soma2=somaPordeducaoUnidadeFuncaoPdt[pdt.Salario_maternidade][pdt.Unidade][pdt.Funcao]
    }


	for _, iadvh := range iadvhs {
		if _, ok := somaPordeducaoUnidadeFuncaoPdt[iadvh.Salario_maternidade]; !ok {
			somaPordeducaoUnidadeFuncaoPdt[iadvh.Salario_maternidade] = make(map[string]map[string]float64)
		}
		if _, ok := somaPordeducaoUnidadeFuncaoPdt[iadvh.Salario_maternidade][iadvh.Unidade]; !ok {
			somaPordeducaoUnidadeFuncaoPdt[iadvh.Salario_maternidade][iadvh.Unidade] = make(map[string]float64)
		}
		somaPordeducaoUnidadeFuncaoPdt[iadvh.Salario_maternidade][iadvh.Unidade][iadvh.Funcao] += iadvh.Salario_maternidade+iadvh.Salario_familia+iadvh.Media_salario_maternidade
        soma1=somaPordeducaoUnidadeFuncaoPdt[iadvh.Salario_maternidade][iadvh.Unidade][iadvh.Funcao]
	}

	


    for i:= range pdts {
        
        pdts[i].Deducao_inssfo = soma1+soma2
    }






    somaPorInsalubridadeUnidadeFuncaoPdt := make(map[float64]map[string]map[string]float64)
		
		
			for _, pdt := range pdts {
				if _, ok := somaPorInsalubridadeUnidadeFuncaoPdt[pdt.Insalubridade_40]; !ok {
					somaPorInsalubridadeUnidadeFuncaoPdt[pdt.Insalubridade_40] = make(map[string]map[string]float64)
				}
				if _, ok := somaPorInsalubridadeUnidadeFuncaoPdt[pdt.Insalubridade_40][pdt.Unidade]; !ok {
					somaPorInsalubridadeUnidadeFuncaoPdt[pdt.Insalubridade_40][pdt.Unidade] = make(map[string]float64)
				}
				somaPorInsalubridadeUnidadeFuncaoPdt[pdt.Insalubridade_40][pdt.Unidade][pdt.Funcao] += pdt.Insalubridade_40 + pdt.Insalubridade_sem_base_40+pdt.Insalubridade_sem_minimo_20+ pdt.Insalubridade_sem_minimo_20+pdt.Insalubridade_retroativa+pdt.Media_insalubridade_prorrogacao_salario_maternidade
			  soma2=somaPorInsalubridadeUnidadeFuncaoPdt[pdt.Insalubridade_40][pdt.Unidade][pdt.Funcao]
            }
		
			
			for _, iadvh := range iadvhs {
				if _, ok := somaPorInsalubridadeUnidadeFuncaoPdt[iadvh.Insalubridade_40]; !ok {
					somaPorInsalubridadeUnidadeFuncaoPdt[iadvh.Insalubridade_40] = make(map[string]map[string]float64)
				}
				if _, ok := somaPorInsalubridadeUnidadeFuncaoPdt[iadvh.Insalubridade_40][iadvh.Unidade]; !ok {
					somaPorInsalubridadeUnidadeFuncaoPdt[iadvh.Insalubridade_40][iadvh.Unidade] = make(map[string]float64)
				}
				somaPorInsalubridadeUnidadeFuncaoPdt[iadvh.Insalubridade_40][iadvh.Unidade][iadvh.Funcao] += iadvh.Diferenca_insalubridade+ iadvh.Insalubridade_20+ iadvh.Insalubridade_40+iadvh.Insalubridade_mes_anterior_40
			   soma1=somaPorInsalubridadeUnidadeFuncaoPdt[iadvh.Insalubridade_40][iadvh.Unidade][iadvh.Funcao]
            }
		
		

            for i:= range pdts {
                
                pdts[i].Insalubridade = soma1+soma2
            }



            somaPorencargosUnidadeFuncaoPdt := make(map[float64]map[string]map[string]float64)


	for _, pdt := range pdts {
		if _, ok := somaPorencargosUnidadeFuncaoPdt[pdt.Aviso_previo_indenizado]; !ok {
			somaPorencargosUnidadeFuncaoPdt[pdt.Aviso_previo_indenizado] = make(map[string]map[string]float64)
		}
		if _, ok := somaPorencargosUnidadeFuncaoPdt[pdt.Aviso_previo_indenizado][pdt.Unidade]; !ok {
			somaPorencargosUnidadeFuncaoPdt[pdt.Aviso_previo_indenizado][pdt.Unidade] = make(map[string]float64)
		}
		somaPorencargosUnidadeFuncaoPdt[pdt.Aviso_previo_indenizado][pdt.Unidade][pdt.Funcao] += (pdt.Fgts_folha+pdt.Ferias+pdt.Ferias_nao_gozadas+pdt.Indenizacao_ferias_nao_gozadas+pdt.Ferias_sobre_aviso_indenizado+pdt.Aviso_previo_indenizado+pdt.Decimo_terceiro_sobre_aviso_indenizado+pdt.Ferias_vencidas+pdt.Ferias_proporcionais+pdt.Fgts_recisao+pdt.Fgts_decimo_terceiro_salario_recisao+pdt.Um_terco_ferias_recisao+pdt.Um_terco_ferias+pdt.Fgts_ferias) - (pdt.Gratificacao_produtor_salario_maternidade+pdt.Gratificacao_dedicacao_exclusiva_salario_maternidade+pdt.Gratificacao_interiorizacao_salario_maternidade+pdt.Salario_maternidade_jovem_aprendiz+pdt.Salario_maternidade+pdt.Media_horas_extras_salario_maternidade+pdt.Media_insalubridade_salario_maternidade+pdt.Gratificacao_sobre_maternidade_supervisor_enfermagem+pdt.Adicional_graduacao_15_tecnico_enfermagem_licenca_maternidade+pdt.Media_gratificacao_sobre_maternidade+pdt.Salario_familia+pdt.Salario_familia_indenizado )
	  soma2=somaPorencargosUnidadeFuncaoPdt[pdt.Aviso_previo_indenizado][pdt.Unidade][pdt.Funcao]
    }

	for _, iadvh := range iadvhs {
		if _, ok := somaPorencargosUnidadeFuncaoPdt[iadvh.Aviso_previo_indenizado]; !ok {
			somaPorencargosUnidadeFuncaoPdt[iadvh.Aviso_previo_indenizado] = make(map[string]map[string]float64)
		}
		if _, ok := somaPorencargosUnidadeFuncaoPdt[iadvh.Aviso_previo_indenizado][iadvh.Unidade]; !ok {
			somaPorencargosUnidadeFuncaoPdt[iadvh.Aviso_previo_indenizado][iadvh.Unidade] = make(map[string]float64)
		}
		somaPorencargosUnidadeFuncaoPdt[iadvh.Aviso_previo_indenizado][iadvh.Unidade][iadvh.Funcao] += (iadvh.Um_terco_ferias_mes+iadvh.Um_terco_ferias_proximo_mes+iadvh.Decimo_terceiro_indenizado+iadvh.Decimo_terceiro_licenca_maternidade+iadvh.Decimo_terceiro_recisao+iadvh.Adicional_um_terco_ferias+iadvh.Adicional_um_terco_ferias_proporcional_recisao+iadvh.Aviso_previo_indenizado+iadvh.Ferias_indenizadas+iadvh.Ferias_mes+iadvh.Ferias_proximo_mes+iadvh.Ferias_proporcionais+iadvh.Media_decimo_terceiro_rescisao+iadvh.Media_aviso_previo+iadvh.Media_ferias+iadvh.Media_ferias_proporcionais+iadvh.Multa_art_479_clt)- ( iadvh.Salario_maternidade+iadvh.Salario_familia+iadvh.Media_salario_maternidade)
	  soma1=somaPorencargosUnidadeFuncaoPdt[iadvh.Aviso_previo_indenizado][iadvh.Unidade][iadvh.Funcao]
    }

    for i:= range pdts{

        pdts[i].Encargosfo = soma1+soma2

    }


    somaPorDsrUnidadeFuncaoPdt := make(map[float64]map[string]map[string]float64)


	for _, pdt := range pdts {
		if _, ok := somaPorDsrUnidadeFuncaoPdt[pdt.Aviso_previo_indenizado]; !ok {
			somaPorDsrUnidadeFuncaoPdt[pdt.Aviso_previo_indenizado] = make(map[string]map[string]float64)
		}
		if _, ok := somaPorDsrUnidadeFuncaoPdt[pdt.Aviso_previo_indenizado][pdt.Unidade]; !ok {
			somaPorDsrUnidadeFuncaoPdt[pdt.Aviso_previo_indenizado][pdt.Unidade] = make(map[string]float64)
		}
		somaPorDsrUnidadeFuncaoPdt[pdt.Aviso_previo_indenizado][pdt.Unidade][pdt.Funcao] += pdt.Dsr_sobre_adicional_noturno + pdt.Dsr_sem_adicional_noturno_mes_anterior
	  soma2=somaPorDsrUnidadeFuncaoPdt[pdt.Aviso_previo_indenizado][pdt.Unidade][pdt.Funcao]
    }

	for _, iadvh := range iadvhs {
		if _, ok := somaPorDsrUnidadeFuncaoPdt[iadvh.Aviso_previo_indenizado]; !ok {
			somaPorDsrUnidadeFuncaoPdt[iadvh.Aviso_previo_indenizado] = make(map[string]map[string]float64)
		}
		if _, ok := somaPorDsrUnidadeFuncaoPdt[iadvh.Aviso_previo_indenizado][iadvh.Unidade]; !ok {
			somaPorDsrUnidadeFuncaoPdt[iadvh.Aviso_previo_indenizado][iadvh.Unidade] = make(map[string]float64)
		} 

		somaPorDsrUnidadeFuncaoPdt[iadvh.Aviso_previo_indenizado][iadvh.Unidade][iadvh.Funcao] += iadvh.Descanso_semanal_remunerado
	   soma1=somaPorDsrUnidadeFuncaoPdt[iadvh.Aviso_previo_indenizado][iadvh.Unidade][iadvh.Funcao]
    }

	
	

	


    for i:= range pdts {
        pdts[i].Dsrfo = soma1+soma2
    }



    somaPortotalsalariofolhaUnidadeFuncaoPdt := make(map[float64]map[string]map[string]float64)
		



    for _, pdt := range pdts {
		if _, ok := somaPortotalsalariofolhaUnidadeFuncaoPdt[pdt.Total_salario_folha]; !ok {
			somaPortotalsalariofolhaUnidadeFuncaoPdt[pdt.Total_salario_folha] = make(map[string]map[string]float64)
		}
		if _, ok := somaPortotalsalariofolhaUnidadeFuncaoPdt[pdt.Total_salario_folha][pdt.Unidade]; !ok {
			somaPortotalsalariofolhaUnidadeFuncaoPdt[pdt.Total_salario_folha][pdt.Unidade] = make(map[string]float64)
		}
        fmt.Println( pdt.Salariobasefo)
        fmt.Println(pdt.Auxilio_alimentacaofo)
        fmt.Println( pdt.Gratificacaofo)
        fmt.Println(pdt.Insalubridade)
        fmt.Println(pdt.Periculosidadefo)
        fmt.Println(pdt.Adicional_noturnofo)
        fmt.Println(pdt.Dsrfo)
		somaPortotalsalariofolhaUnidadeFuncaoPdt[pdt.Total_salario_folha][pdt.Unidade][pdt.Funcao] += pdt.Salariobasefo + pdt.Auxilio_alimentacaofo+pdt.Gratificacaofo+pdt.Insalubridade+pdt.Periculosidadefo+pdt.Adicional_noturnofo+pdt.Dsrfo
          soma2=somaPortotalsalariofolhaUnidadeFuncaoPdt[pdt.Total_salario_folha][pdt.Unidade][pdt.Funcao]
        fmt.Println("valor do soma salario1 %0.2f", somaPortotalsalariofolhaUnidadeFuncaoPdt[pdt.Total_salario_folha][pdt.Unidade][pdt.Funcao])
        break
	}

	for _, iadvh := range iadvhs {
		if _, ok := somaPortotalsalariofolhaUnidadeFuncaoPdt[iadvh.Total_salario_folha]; !ok {
			somaPortotalsalariofolhaUnidadeFuncaoPdt[iadvh.Total_salario_folha] = make(map[string]map[string]float64)
		}
		if _, ok := somaPortotalsalariofolhaUnidadeFuncaoPdt[iadvh.Total_salario_folha][iadvh.Unidade]; !ok {
			somaPortotalsalariofolhaUnidadeFuncaoPdt[iadvh.Total_salario_folha][iadvh.Unidade] = make(map[string]float64)
		} 
        somaPortotalsalariofolhaUnidadeFuncaoPdt[iadvh.Total_salario_folha][iadvh.Unidade][iadvh.Funcao] += iadvh.Salario_base+iadvh.Gratificaoia+iadvh.Insalubridade+iadvh.Periculosidadeia+iadvh.Adicional_noturnoia+iadvh.Dsria
        fmt.Println(iadvh.Salario_base)
     //   fmt.Println(iadvh.Auxilio_alimentacao)
        fmt.Println(iadvh.Gratificacao)
        fmt.Println(iadvh.Insalubridade)
        fmt.Println(iadvh.Periculosidade)
        fmt.Println(iadvh.Adicional_noturno)
        fmt.Println(iadvh.Dsria)
        fmt.Println("valor do soma salario2 %0.2f", somaPortotalsalariofolhaUnidadeFuncaoPdt[iadvh.Total_salario_folha][iadvh.Unidade][iadvh.Funcao])
        break
        soma1= somaPortotalsalariofolhaUnidadeFuncaoPdt[iadvh.Total_salario_folha][iadvh.Unidade][iadvh.Funcao]
	}
		
			
			
			
	
			
            for i:= range pdts {
                
                pdts[i].Total_salario_folha = soma1+soma2
            }




            somaPortotalsalariomensalfolhaUnidadeFuncaoPdt := make(map[float64]map[string]map[string]float64)
		
		
			for _, pdt := range pdts {
				if _, ok := somaPortotalsalariomensalfolhaUnidadeFuncaoPdt[pdt.Total_mensal_folha]; !ok {
					somaPortotalsalariomensalfolhaUnidadeFuncaoPdt[pdt.Total_mensal_folha] = make(map[string]map[string]float64)
				}
				if _, ok := somaPortotalsalariomensalfolhaUnidadeFuncaoPdt[pdt.Total_mensal_folha][pdt.Unidade]; !ok {
					somaPortotalsalariomensalfolhaUnidadeFuncaoPdt[pdt.Total_mensal_folha][pdt.Unidade] = make(map[string]float64)
				}

				//soma1:=somaPorencargosUnidadeFuncaoPdt 
				//soma3:= somaPortotalsalariofolhaUnidadeFuncaoPdt
				somaPortotalsalariomensalfolhaUnidadeFuncaoPdt[pdt.Total_mensal_folha][pdt.Unidade][pdt.Funcao] += pdt.Total_salario_folha + pdt.Encargosfo
                soma2=somaPortotalsalariomensalfolhaUnidadeFuncaoPdt[pdt.Total_mensal_folha][pdt.Unidade][pdt.Funcao]
			}
		
			
			for _, iadvh := range iadvhs {
				if _, ok := somaPortotalsalariomensalfolhaUnidadeFuncaoPdt[iadvh.Total_mensal_folha]; !ok {
					somaPortotalsalariomensalfolhaUnidadeFuncaoPdt[iadvh.Total_mensal_folha] = make(map[string]map[string]float64)
				}
				if _, ok := somaPortotalsalariomensalfolhaUnidadeFuncaoPdt[iadvh.Total_mensal_folha][iadvh.Unidade]; !ok {
					somaPortotalsalariomensalfolhaUnidadeFuncaoPdt[iadvh.Total_mensal_folha][iadvh.Unidade] = make(map[string]float64)
				}

				
				somaPortotalsalariomensalfolhaUnidadeFuncaoPdt[iadvh.Total_mensal_folha][iadvh.Unidade][iadvh.Funcao] += iadvh.Encargosia + iadvh.Total_salario_folha
                soma1 =somaPortotalsalariomensalfolhaUnidadeFuncaoPdt[iadvh.Total_mensal_folha][iadvh.Unidade][iadvh.Funcao]
			}
		
		

            for i:= range pdts {
                
                pdts[i].Total_mensal_folha = soma1+soma2
            }



            somaPortotalsalarioanualfolhaUnidadeFuncaoPdt := make(map[float64]map[string]map[string]float64)
		
		
for _, pdt := range pdts {
	if _, ok := somaPortotalsalarioanualfolhaUnidadeFuncaoPdt[pdt.Total_anual_folha]; !ok {
		somaPortotalsalarioanualfolhaUnidadeFuncaoPdt[pdt.Total_anual_folha] = make(map[string]map[string]float64)
	}
	if _, ok := somaPortotalsalarioanualfolhaUnidadeFuncaoPdt[pdt.Total_anual_folha][pdt.Unidade]; !ok {
		somaPortotalsalarioanualfolhaUnidadeFuncaoPdt[pdt.Total_anual_folha][pdt.Unidade] = make(map[string]float64)
	}
	somaPortotalsalarioanualfolhaUnidadeFuncaoPdt[pdt.Total_anual_folha][pdt.Unidade][pdt.Funcao] += (pdt.Total_mensal_folha)*12
    soma2=somaPortotalsalarioanualfolhaUnidadeFuncaoPdt[pdt.Total_anual_folha][pdt.Unidade][pdt.Funcao]
}


for _, iadvh := range iadvhs {
	if _, ok := somaPortotalsalarioanualfolhaUnidadeFuncaoPdt[iadvh.Total_anual_folha]; !ok {
		somaPortotalsalarioanualfolhaUnidadeFuncaoPdt[iadvh.Total_anual_folha] = make(map[string]map[string]float64)
	}
	if _, ok := somaPortotalsalarioanualfolhaUnidadeFuncaoPdt[iadvh.Total_anual_folha][iadvh.Unidade]; !ok {
		somaPortotalsalarioanualfolhaUnidadeFuncaoPdt[iadvh.Total_anual_folha][iadvh.Unidade] = make(map[string]float64)
	}
	somaPortotalsalarioanualfolhaUnidadeFuncaoPdt[iadvh.Total_anual_folha][iadvh.Unidade][iadvh.Funcao] += (iadvh.Total_mensal_folha)*12
    soma1= somaPortotalsalarioanualfolhaUnidadeFuncaoPdt[iadvh.Total_anual_folha][iadvh.Unidade][iadvh.Funcao]
}



for i:= range pdts{
    pdts[i].Total_anual_folha= soma1+soma2
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
       go worker(&wg, done,Calculo)
   }

   // Aguarda a conclusão de todas as goroutines
   go func() {
       wg.Wait()
       close(done)
   }()

   // Aguarda o canal done ser fechado
   <-done



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