package models


import (
	//"gorm.io/gorm"
	"time"

)

type PDT struct {

	//gorm.Model

	
	      ID int  `json:"-"`
		Id_funcionario int `json:"-"`
		Nome_funcionario string `json:"-"`
		Unidade string `json:"unidade"`
		Funcao string  `json:"setor"`
		Cpf string `json:"-"`
		Data_admissao time.Time `json:"-"`
		Data_demissao time.Time `json:"-"`
		Salario_base float64 `json:"-"`
		Qtde_salario_base int `json:"-"`
		Jovem_aprendiz float64 `json:"-"`
		Qtde_jovem_aprendiz int `json:"-"`
		Saldo_salario float64 `json:"-"`
		Qtde_saldo_salario float64 `json:"-"`
		Saldo_salario_horista float64 `json:"-"`
		Qtde_saldo_salario_horista float64 `json:"-"`
		Ferias float64 `json:"-"`
		Periculosidade float64 `json:"-"`
		Auxilio_alimentacao float64 `json:"-"`
		Insalubridade_sem_base_40 float64 `json:"-"`
		Auxilio_alimentacao_retroativo float64 `json:"-"`
		Gratificacao_funcao_retroativo float64 `json:"-"`
		Diferenca_salarial float64 `json:"-"`
		Adicional_noturno_20 float64 `json:"-"`
		Qtde_adicional_noturno_20 int `json:"-"`
		Adicional_noturno_30 float64 `json:"-"`
		Qtde_adicional_noturno_30 int `json:"-"`
		Insalubridade_40 float64 `json:"-"`
		Insalubridade_sem_minimo_20 float64 `json:"-"`
		Insalubridade_sem_minimo_40 float64 `json:"-"`
		Gratificacao_funcao float64 `json:"-"`


		
			
		Devolucao_falta float64 `json:"-"`
		Gratificacao_dedicacao_exclusiva float64 `json:"-"`
		Gratificacao_produtividade float64 `json:"-"`
		Reembolso_faltas float64 `json:"-"`
		Adicional_noturno_20_valor float64 `json:"-"`
		Ferias_nao_gozadas float64 `json:"-"`
		Adicional_noturno_30_mes_anterior float64 `json:"-"`
		Qtde_adicional_noturno_30_mes_anterior int `json:"-"`
		Dsr_sem_adicional_noturno_mes_anterior float64 `json:"-"`
		Diferenca_adicional_noturno float64 `json:"-"`
		Hora_extra_50_retroativo_01 int `json:"-"`
		Qtde_hora_extra_50_retroativo_01 float64 `json:"-"`
		Salario_substituicao float64 `json:"-"`
		Salario_substituicao_mes_anterior float64 `json:"-"`
		Adicional_noturno_20_retroativo float64 `json:"-"`
		Qtde_adicional_noturno_20_retroativo int `json:"-"`
		Gratificacao_funcao_supervisor_enfermagem float64 `json:"-"`
		Dsr_sem_horas_extras_retroativas float64 `json:"-"`
		Adicional_graduacao_15_tecnico_enfermage float64 `json:"-"`
	//parte2
		Gratificacao_5_sobre_salario_base float64 `json:"-"`
		Estorno_desconto_vale_transporte float64 `json:"-"`
		Prorrogacao_media_deducao_exclusiva_salario_maternidade float64 `json:"-"`
		Gratificacao_funcao_instrumentacao_cirurgico float64 `json:"-"`
		Gratificacao_por_interiorizacao float64 `json:"-"`
		Suspensao float64 `json:"-"`
		Gratificacao_produtor_salario_maternidade float64 `json:"-"`
		Gratificacao_dedicacao_exclusiva_salario_maternidade float64 `json:"-"`
		Gratificacao_interiorizacao_salario_maternidade float64 `json:"-"`
		Dias_atestado float64 `json:"-"`
		Qtde_dias_atestado float64 `json:"-"`
		Salario_maternidade_jovem_aprendiz float64 `json:"-"`
		Qtde_salario_maternidade_jovem_aprendiz int `json:"-"`
		Indenizacao_ferias_nao_gozadas float64 `json:"-"`
		Gratificacao float64 `json:"-"`
		Insalubridade_retroativa float64 `json:"-"`
		Salario_maternidade float64 `json:"-"`
		Qtde_salario_maternidade int `json:"-"`
		Media_horas_extras_salario_maternidade float64 `json:"-"`
		Media_insalubridade_salario_maternidade float64 `json:"-"`
		Media_adicional_noturno_salario_maternidade float64 `json:"-"`
		Prorrogacao_licenca_maternidade float64 `json:"-"`
		Media_horas_extras_prorrogacao_salario_maternidade float64 `json:"-"`
		Media_insalubridade_prorrogacao_salario_maternidade float64 `json:"-"`
		Media_adicional_noturno_prorrogacao_salario_maternidade float64 `json:"-"`
		Gratificacao_sobre_maternidade_supervisor_enfermagem float64 `json:"-"`
		Adicional_graduacao_15_tecnico_enfermagem_licenca_maternidade float64 `json:"-"` //aqui
		Media_gratificacao_sobre_maternidade float64 `json:"-"`
		//Media_gratificacao_sobre_maternidade *float64 `json:"media_gratificacao_sobre_maternidade"`
		Adicional_gratificacao_15_tecnico_enfermagem_prorrogacao_licenca_maternidade float64 `json:"-"`
		Seguro_acidente float64 `json:"-"`
		Desconto_dsr_em_dias float64 `json:"-"`
		Qtde_desconto_dr_em_dias float64 `json:"-"`
		Atrasos float64 `json:"-"`
		Qtde_atrasos int `json:"-"`
		Faltas_em_dias float64 `json:"-"`
		Qtde_faltas_em_dias int `json:"-"`
		Hora_extra_50 float64 `json:"-"`
		Qtde_hora_extra_50 int `json:"-"`
		Hora_extra_100 float64 `json:"-"`
		Qtde_hora_extra_100 int `json:"-"`
		Hora_extra_sobre_adicionais_50 float64 `json:"-"`
		Qtde_hora_extra_sobre_adicionais_50 int `json:"-"`
		Hora_extra_sobre_adicionais_60 float64 `json:"-"`
		Qtde_hora_extra_sobre_adicionais_60 int `json:"-"`
		Hora_extra_sobre_adicionais_100 float64 `json:"-"`
		Qtde_hora_extra_sobre_adicionais_100 int `json:"-"`
		Pensao_alimenticia_sobre_deducoes float64 `json:"-"`
		Hora_extra_retroativa_60 float64 `json:"-"`
		Qtde_hora_extra_retroativa_60 int `json:"-"`
		Hora_extra_mes_anterior_50 float64 `json:"-"`
		Qtde_hora_extra_mes_anterior_50 int `json:"-"`
		Hora_extra_50_valor float64 `json:"-"`
		Hora_extra_100_mes_anterior float64 `json:"-"`
		Qtde_hora_extra_100_mes_anterior int `json:"-"`
		Hora_extra_50_retroativo float64 `json:"-"`
		Qtde_hora_extra_50_retroativo int `json:"-"`
		Dsr_sobre_horas_extras float64 `json:"-"`
		Dsr_sobre_adicional_noturno float64 `json:"-"`
		Hora_extra_100_valor float64 `json:"-"`
		Hora_extra_60_valor float64 `json:"-"`
		Licenca_remunerada float64 `json:"-"`
		Ferias_sobre_aviso_indenizado float64 `json:"-"`
		Qtde_ferias_sobre_aviso_indenizado int `json:"-"`
		Faltas_ferias_proporcionais float64 `json:"-"`
		Aviso_previo_indenizado float64 `json:"-"`
		Qtde_aviso_previo_indenizado int `json:"-"`
		Decimo_terceiro_proporcional float64 `json:"-"`
		Qtde_decimo_terceiro_proporcional float64 `json:"-"`
		Decimo_terceiro_sobre_aviso_indenizado float64 `json:"-"`
		Qtde_decimo_terceiro_sobre_aviso_indenizado float64 `json:"-"`
		Salario_familia_indenizado float64 `json:"-"`
		Ferias_vencidas float64 `json:"-"`
		Qtde_ferias_vencidas int `json:"-"`
		Ferias_proporcionais float64 `json:"-"`
		Qtde_ferias_proporcionais int `json:"-"`
		Fgts_recisao float64 `json:"-"`
		Fgts_decimo_terceiro_salario_recisao float64 `json:"-"`
		Um_terco_ferias_recisao float64 `json:"-"`
		Multa_fgts float64  `json:"-"`
		Um_terco_ferias float64 `json:"-"`
		Salario_familia float64 `json:"-"`
		Mensalidade_sindical float64 `json:"-"`
		Primeira_parcela_decimo_terceirom float64 `json:"-"`
		Insuficiencia_saldo_desconto float64 `json:"-"`
		Aviso_previo_indenizado_descontado float64 `json:"-"`
		Auxilio_alimentacao_desconto float64 `json:"-"`
		Vale_transporte_desconto float64 `json:"-"`
		Pensao_alimenticia_sobre_minimo float64 `json:"-"`
		Desconto_adiantamento float64 `json:"-"`
		Pensao_30_sobre_salario_minimo float64 `json:"-"`
		Adicional_noturno_30_valor float64 `json:"-"`
		Pensao_alimenticia_18 float64 `json:"-"`
		Pensao_alimenticia_22 float64 `json:"-"`
		Pensao_alimenticia_sobre_minimo_20 float64 `json:"-"`
		Pensao_alimenticia float64 `json:"-"`
		//Pensao_alimenticia *float64 `json:"pensao_alimenticia"`
		Faltas_valor float64 `json:"-"`
		Pensao_alimenticia_sobre_bruto_15 float64 `json:"-"`
		Desconto_funben_titular float64 `json:"-"`
		Desconto_funben_dependente float64 `json:"-"`
		Media_gratificacao_funcao_prorrogacao_maternidade float64 `json:"-"`
		Inss_rescisao float64 `json:"-"`
		Inss_decimo_terceiro_rescisao float64 `json:"-"`
		Inss_ferias float64 `json:"-"`
		Inss_folha float64 `json:"-"`
		Irrf_rescisao float64 `json:"-"`
		Irrf_decimo_terceiro_rescisao float64 `json:"-"`
		Irrf_ferias float64 `json:"-"`
		Irrf_folha float64 `json:"-"`
		Fgts_ferias float64 `json:"-"`
		Fgts_folha float64 `json:"-"`
		Pensao_alimenticia_0_3 float64 `json:"-"`
		Ajuste_demonstrativo_esocial float64 `json:"-"`
		Pensao_sobre_bruto_25 float64 `json:"-"`
		Gratificacao_funcao_supervisor_enfermagem_sl float64 `json:"-"`
		Reembolso_educacional float64 `json:"-"`
		Ajuste_demonstrativo_rra_esocial float64 `json:"-"`
		Assistencia_financeira_complementar_05_2 float64 `json:"-"`
		Assistencia_financeira_complementar_06_2 float64 `json:"-"`
		Assistencia_financeira_complementar_07_2 float64 `json:"-"`
		Assistencia_financeira_complementar_08_2 float64 `json:"-"`
		Assistencia_financeira_complementar_09_2 float64 `json:"-"`
		Insuficiencia_saldo_provento float64 `json:"-"`
     //  Carga_horaria int `json:"carga_horaria"`
	
		Rh int `json:"rh"`	 
	    Salariobasefo float64 `json:"salario_folha"`
		Auxilio_alimentacaofo float64 `json:"auxilio_alimentacao"`
		Gratificacaofo float64 `json:"gratificacao_folha"`
		Insalubridade float64 `json:"insalubridade_folha"`
		Periculosidadefo float64 `json:"periculosidade_folha"`
		Adicional_noturnofo float64 `json:"adicional_noturno_folha"`
	    //Adicional_noturnoo int `json:"adicional_noturno_folha"`
        Dsrfo float64 `json:"dsr_folha"`

		Total_salario_folha float64 `json:"total_salario_folha`
		Encargosfo float64 `json:"encargos_folha"`
		//vale transporte
		Total_mensal_folha float64 `json:"total_mensal_folha`
		Total_anual_folha float64 `json:"total_anual_folha`
		Hora_extrafo float64 `json:"hora_extras"`
		
		Deducao_inssfo float64 `json:"-"` //diminuir com encargos
		
		

			
		
	
}
	
	
