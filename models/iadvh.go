package models

import (
	"gorm.io/gorm"
	"time"
	//"database/sql"
)



type Iadvh struct{

	gorm.Model
Rh int `json:"rh"`	
	Id_funcionario int `json:"id_funcionario"`//
	Nome_funcionario string `json:"nome_funcionario"`//
	Funcao string  `json:"funcao"`//
	Cpf string `json:"cpf"`//
	Unidade string `json:"unidade"`//
	Data_admissao time.Time `json:"data_admissao"`//
	Data_demissao time.Time `json:"data_demissao"`//
	Um_terco_ferias_mes float64 `json:"um_terco_ferias_mes`//
	Um_terco_ferias_proximo_mes float64 `json:um_terco_ferias_proximo_mes`//
	Um_terco_ferias_pagas_mes_anterior float64 `json:um_terco_ferias_pagas_mes_anterior`//
	Decimo_terceiro_indenizado float64 `json:decimo_terceiro_indenizado`//
	Decimo_terceiro_licenca_maternidade float64 `json:decimo_terceiro_licenca_maternidade`//
	Decimo_terceiro_recisao float64 `json:decimo_terceiro_recisao`//
	Adiantamento_decimo_terceiro_salario float64 `json:adiantamento_decimo_terceiro_salario`//
	Adiantamento_conforme_recebido float64 `json:adiantamento_conforme_recebido`//
	Adicional_um_terco_ferias float64 `json:adicional_um_terco_ferias`//
	Adicional_um_terco_ferias_proporcional_recisao float64 `json:adicional_um_terco_ferias_proporcional_recisao`//
	Adicional_noturno float64 `json:adicional_noturno`//
	Adicional_noturno_mes_anterior float64 `json:adicional_noturno_mes_anterior`//
	Ajuste_base_fgts_decimo_terceiro float64 `json:ajuste_base_fgts_decimo_terceiro`//
	Ajuste_base_inss float64 `json:"ajuste_base_inss"`//
	Ajuste_base_inss_mes_anterior float64 `json:"ajuste_base_inss_mes_anterior"`//
	Assistencia_financeira_complementar_05_2023 float64 `json:"assistencia_financeira_complementar_05_2023"`//
	Assistencia_financeira_complementar_06_2023 float64 `json:"assistencia_financeira_complementar_06_2023"`//
	Assistencia_financeira_07_2023 float64 `json:"assistencia_financeira_07_2023"`//
	Assistencia_financeira_08_2023 float64 `json:"assistencia_financeira_08_2023"`//
	Assistencia_financeira_09_2023 float64 `json:"assistencia_financeira_09_2023"`//
	Atestado_medico float64 `json:"atestado_medico"`//
	Atrasos float64 `json:"atrasos"`//
	Auxilio_creche float64 `json:"auxilio_creche"`//
	Auxilio_funeral float64 `json:"auxilio_funeral"`//
	Aviso_previo_indenizado float64 `json:"aviso_previo_indenizado"`//
	Base_calculo_inss_outro_emprego float64 `json:"base_calculo_inss_outro_emprego"`//
	Base_irrf_ferias_envelope float64 `json:"base_irrf_ferias_envelope"`//
	Base_inss_ferias_mes_outro_emprego_mes_anterior float64 `json:"base_inss_ferias_mes_outro_emprego_mes_anterior"`//
	Base_liquida_irrf float64 `json:"base_liquida_irrf"`//
	Base_liquida_irrf_decimo_terceiro float64 `json"base_liquida_irrf_decimo_terceiro"`//
	Base_liquida_irrf_ferias float64 `json:"base_liquida_irrf_ferias"`//
	Base_pis float64 `json:"base_pis"`//
	Desconto_pagamento_indevido_com_incidencia float64 `json:"desconto_pagamento_indevido_com_incidencia"`//
	//sem incidencia
	Descanso_semanal_remunerado float64 `json:"descanso_semanal_remunerado"`//
	Descanso_semanal_remunerado_perdido float64 `json:"descanso_semanal_remunerado_perdido"`//
	Desconto_aviso_previo float64 `json:"desconto_aviso_previo"`//
	Desconto_judicial_30_rendimentos float64 `json:"desconto_judicial_30_rendimentos"`//
	Descontos_diversos float64 `json:"descontos_diversos"`//
	Devolucao_desconto_indevido float64 `json:"devolucao_desconto_indevido"`//
	Devolucao_faltas float64 `json:"devolucao_faltas"`//
	Desconto_plano_saude_hapvida float64 `json:"desconto_plano_saude_hapvida"`//
	Desconto_plano_odontologico_integral float64 `json:"desconto_plano_odontologico_integral"`//
	Diarias float64 `json:"diarias"`//
	Dias_trabalhados float64 `json:"dias_trabalhados"`//
	Diferenca_insalubridade float64 `json:"diferenca_insalubridade"`//
	Diferenca_descricao_plano_saude_mes_anterior float64 `json:"diferenca_descricao_plano_saude_mes_anterior"`//
	Diferenca_salarial float64 `json:"diferenca_salarial"`//
	Dsr_suspensao float64 `json:"dsr_suspensao"`//
	Estorno_faltas float64 `json:"estorno_faltas"`//
	Falta_dsr float64 `json:"falta_dsr"`//
	Falta_dsr_escala float64 `json:"falta_dsr_escala"`//
	Falta_dsr_escala_mes_rescisao float64 `json:"falta_dsr_escala_mes_rescisao"`//
	Falta_dsr_mes_rescisao float64 `json:"falta_dsr_mes_rescisao"`//
	Falta_dsr_mes_anterior_com_incidencia float64 `json:"falta_dsr_mes_anterior_com_incidencia"`//
	Falta_mes_rescisao float64 `json:"falta_mes_rescisao"`//
	Faltas float64 `json:"faltas"`//
	Faltas_dsr_mes_anterior_rescisao float64 `json:"faltas_dsr_mes_anterior_rescisao"`
	Faltas_mes_anterior float64 `json:"faltas_mes_anterior"`//
	Faltas_mes_anterior_com_incidencia float64 `json:"faltas_mes_anterior_com_incidencia"`//
	Ferias_indenizadas float64 `json:"ferias_indenizadas"`//
	Ferias_mes float64 `json:"ferias_mes"`//
	Ferias_proximo_mes float64 `json:"ferias_proximo_mes"`//
	Ferias_pagas_mes_anterior float64 `json:"ferias_pagas_mes_anterior"`//
	Ferias_proporcionais float64 `json:"ferias_proporcionais"`//
	Ferias_proporcionais_sem_aviso_previo float64 `json:"ferias_proporcionais_sem_aviso_previo"`//
	Ferias_proporcionais_sobre_aviso_previo_indenizado float64 `json:"ferias_proporcionais_sobre_aviso_previo_indenizado"`//
	Ferias_vencidas_sem_aviso_previo float64 `json:"ferias_vencidas_sem_aviso_previo"`//
	Fgts_decimo_terceiro_depositado_rescisao float64 `json:"fgts_decimo_terceiro_depositado_rescisao"`//
	Fgts_decimo_terceiro_rescisao float64 `json:"fgts_decimo_terceiro_rescisao"`//
	Fgts_artigo_22 float64 `json:"fgts_artigo_22"`//
	Fgts_depositado_rescisao float64 `json:"fgts_depositado_rescisao"`//
	Fgts_quitacao float64 `json:"fgts_quitacao"`//
	Gratificacao float64 `json:"gratificacao"`//
	Gratificacao_cargo_confianca float64 `json:"gratificacao_cargo_confianca"`//
	Gratificacao_indenizatoria float64 `json:"gratificacao_indenizatoria"`//
	Gratificacao_funcao_instrumentacao_cirurgico float64 `json:"gratificacao_instrumentacao_cirurgica"`//
	Gratificacao_mes_anterior float64 `json:"gratificacao_mes_anterior"`//
	Hora_extra_100 float64 `json:"hora_extra_100"`//
	Hora_extra_50 float64 `json:"hora_extra_50"`//
	Hora_extra_60 float64 `json:"hora_extra_60"`//
	Hora_extra_60_mes_anterior float64 `json:"hora_extra_60_mes_anterior"`//
	Horas_extras_100_mes_anterior float64 `json:"horas_extras_100_mes_anterior"`//
	Horas_extras_50_mes_anterior float64 `json:"horas_extras_50_mes_anterior"`//
	Horas_falta_dsr float64 `json:"horas_falta_dsr"`//
	Horas_noturnas_reduzidas float64 `json:"horas_noturnas_reduzidas"`//
	Indenizacao_art_480 float64 `json:"indenizacao_art_480"`//
	Insalubridade_20 float64 `json:"insalubridade_20"`//
	Insalubridade_40 float64 `json:"insalubridade_40"`//
	Insalubridade_mes_anterior_40 float64 `json:"insalubridade_mes_anterior_40"`//
	Inss float64 `json:"inss"`
	Inss_decimo_terceiro_aliquota_normal float64 `json:"inss"`//
	//inss 13 com aliquota normal
	Inss_decimo_terceiro_salario float64 `json:"inss_decimo_terceiro_salario"`//
	Inss_aliquota_normal float64 `json:"inss_aliquota_normal"`//
	Inss_ferias_mes_anterior_outro_emprego float64  `json:"inss_ferias_mes_anterior_outro_emprego"`//
	Inss_ferias float64 `json:"inss_ferias"`//
	Inss_ferias_aliquota_normal float64 `json:"inss_ferias_aliquota_normal"`//
	Inss_ferias_mes_anterior float64 `json:"inss_ferias_mes_anterior"`//
	Inss_ferias_mes_anterior_aliquota_normal float64 `json:"inss_ferias_mes_anterior_aliquota_normal"`//
	Inss_ferias_proximo_mes float64 `json:"inss_ferias_proximo_mes"`//
	Inss_ferias_proximo_mes_aliquota_normal float64 `json:"inss_ferias_proximo_mes_aliquota_normal"`//
	Inss_outro_emprego float64 `json:"inss_outro_emprego"`//
	Insuficiencia_saldo_desconto float64 `json:"insuficiencia_saldo_desconto"`//
	Insuficiencia_saldo_proventos float64 `json:"insuficiencia_saldo_proventos"`//
	Irrf float64 `json:"irrf"`//
	Irrf_decimo_terceiro_salario float64 `json:"irrf_decimo_terceiro_salario"`//
	Irrf_ferias float64 `json:"irrf_ferias"`//
	Licenca_paternidade float64 `json:"licenca_paternidade"`//
	Liquido_folha float64 `json:"liquido_folha"`//
	Liquido_rescisao float64 `json:"liquido_rescisao"`//
	Media_decimo_terceiro_rescisao float64 `json:"media_decimo_terceiro_rescisao"`//
	Media_aviso_previo float64 `json:"media_aviso_previo"`//
	Media_ferias float64 `json:"media_ferias"`//
	Media_ferias_proporcionais float64 `json:"media_ferias_proporcionais"`//
	Media_salario_maternidade float64  `json:"media_salario_maternidade"`//
	Mensalidade_sindicato_enfermeiros_ma float64 `json:"mensalidade_sindicato_enfermeiros_ma"`//
	Mensalidade_sindicato_enfermeiros_ma_ferias float64 `json:"mensalidade_sindicato_enfermeiros_ma_ferias"`//
	Mensalidade_sindicato_sind_saude float64 `json:"mensalidade_sindicato_sind_saude"`//
	Mensalidade_sindicato_sind_saude_ferias float64 `json:"mensalidade_sindicato_sind_saude_ferias"`//
	Mensalidade_sindicato_sintaema float64 `json:"mensalidade_sindicato_sintaema"`//
	Mensalidade_sindicato_sintaema_ferias float64 `json:"mensalidade_sindicato_sintaema_ferias"`//
	Mensalidade_sindicato_imperatriz float64 `json:"mensalidade_sindicato_imperatriz"`//
	Mensalidade_sindicato_imperatriz_ferias float64 `json:"mensalidade_sindicato_imperatriz_ferias"`//
	Multa_art_479_clt float64 `json:"multa_art_479_clt"`//
	Outros_proventos float64 `json:"outros_proventos"`//
	Pensao_alimenticia_folha float64 `json:"pensao_alimenticia_folha"`//
	Pensao_alimenticia_ferias float64 `json:"pensao_alimenticia_ferias"`//
	Periculosidade float64 `json:"periculosidade"`//
	Pis_1 float64 `json:"pis_1"`//
	Plano_odontologico_integral float64 `json:"plano_odontologico_integral"`//
	Plano_odontologico_premium float64 `json:"plano_odontologico_premium"`//
	Plano_saude_apartamento_hapvida float64 `json:"plano_saude_apartamento_hapvida"`//
	Plano_saude_enfermaria_hapvida float64 `json:"plano_saude_enfermaria_hapvida"`//
	Plano_saude_taxa_adesao_hapvida float64 `json:"plano_saude_taxa_adesao_hapvida"`//
	Plano_saude_apartamento_ferias_hapvida float64 `json:"plano_saude_apartamento_ferias_hapvida"`//
	Plano_saude_enfermaria_ferias_hapvida float64 `json:"plano_saude_enfermaria_ferias_hapvida"`//
	//plano odontologico integral
	Plano_odontologico_premium_ferias float64 `json:"plano_odontologico_premium_ferias"`//
	Salario_familia  float64  `json:"salario_familia"`//
	Salario_maternidade float64  `json:"salario_maternidade"`//
	Saldo_salario float64 `json:"saldo_salario"`//
	Saldo_fgts_banco float64 `json:"saldo_fgts_banco"`//
	Suspensao float64 `json:"suspensao"`//
	Suspensao_ponto float64 `json:"suspensao_ponto"`//
	Ticket_alimentacao float64 `json:"ticket_alimentacao"`//
	Ticket_alimentacao_mes_anterior float64 `json:"ticket_alimentacao_mes_anterior"`//
	Total_calculado_compra_vale_transporte float64 `json:"total_calculado_compra_vale_transporte"`//
	Total_entregue_vale_transporte float64 `json:"total_entregue_vale_transporte"`//
	Vale_transporte float64 `json:"vale_transporte"`//
	//Liquido_folha float64 `json:"liquido_folha"` //tem que tá em outro lugar
	Adicional_noturnoo int `json:"adicional_noturnoo"`
	Inssia  float64 `json:"deducaoinssia"`
	Dsria float64 `json:"dsria"`
	Encargosia float64 `json:"encargosia"`
	Gratificaoia float64  `json:"gratificacaofi`
	Hora_extraia float64 `json:"hora_extraia"`
	Insalubridade float64 `json:"insalubridadeia"`
	Periculosidadeia float64  `json:"periculosidadeia"`
	Salario_base float64 `json:"salarioia_base"`
	Adicional_noturnoia float64 `json:"adicional_noturnoia"`
	
	Total_salario_folha float64 `json:"total_salario_folha`
	//vale transporte
	Total_mensal_folha float64 `json:"total_mensal_folha`
	Total_anual_folha float64 `json:"total_anual_folha`



}