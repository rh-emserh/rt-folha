


package controller

import (
	"example.com/gogin/config"
	"example.com/gogin/models"
	//"gorm.io/gorm"
)

func CalcSalariobase() {
	//var soma float64 = 0.0
	var pdts []models.PDT
	config.DB.Find(&pdts)

	var iadvhs []models.Iadvh
	config.DB.Find(&iadvhs)

	somaPorsalariobaseUnidadeFuncaoPdt := make(map[float64]map[string]map[string]float64)

	for _, pdt := range pdts {
		
		if _, ok := somaPorsalariobaseUnidadeFuncaoPdt[pdt.Saldo_salario]; !ok {
			somaPorsalariobaseUnidadeFuncaoPdt[pdt.Saldo_salario] = make(map[string]map[string]float64)
		}
		if _, ok := somaPorsalariobaseUnidadeFuncaoPdt[pdt.Saldo_salario][pdt.Unidade]; !ok {
			somaPorsalariobaseUnidadeFuncaoPdt[pdt.Saldo_salario][pdt.Unidade] = make(map[string]float64)
		}

		somaPorsalariobaseUnidadeFuncaoPdt[pdt.Saldo_salario][pdt.Unidade][pdt.Funcao] += pdt.Salario_base+pdt.Jovem_aprendiz+pdt.Diferenca_salarial+pdt.Devolucao_falta+pdt.Reembolso_faltas+pdt.Salario_substituicao+pdt.Salario_substituicao_mes_anterior+pdt.Estorno_desconto_vale_transporte+pdt.Dias_atestado+pdt.Prorrogacao_licenca_maternidade+pdt.Licenca_remunerada+pdt.Reembolso_educacional+pdt.Insuficiencia_saldo_provento+pdt.Saldo_salario+pdt.Saldo_salario_horista
       // soma2 = somaPorsalariobaseUnidadeFuncaoPdt[pdt.Saldo_salario][pdt.Unidade][pdt.Funcao]
	}


	for _, iadvh := range iadvhs {
		
		if _, ok := somaPorsalariobaseUnidadeFuncaoPdt[iadvh.Saldo_salario]; !ok {
			somaPorsalariobaseUnidadeFuncaoPdt[iadvh.Saldo_salario] = make(map[string]map[string]float64)
		}
		if _, ok := somaPorsalariobaseUnidadeFuncaoPdt[iadvh.Saldo_salario][iadvh.Unidade]; !ok {
			somaPorsalariobaseUnidadeFuncaoPdt[iadvh.Saldo_salario][iadvh.Unidade] = make(map[string]float64)
		}
	
		somaPorsalariobaseUnidadeFuncaoPdt[iadvh.Saldo_salario][iadvh.Unidade][iadvh.Funcao] += iadvh.Atestado_medico+iadvh.Devolucao_desconto_indevido+iadvh.Devolucao_faltas+iadvh.Desconto_plano_saude_hapvida+iadvh.Desconto_plano_odontologico_integral+iadvh.Dias_trabalhados+iadvh.Diferenca_salarial+iadvh.Horas_noturnas_reduzidas+iadvh.Insuficiencia_saldo_proventos+iadvh.Licenca_paternidade+iadvh.Saldo_salario
       // soma1 = somaPorsalariobaseUnidadeFuncaoPdt[iadvh.Saldo_salario][iadvh.Unidade][iadvh.Funcao]
	}
	// Iniciar uma transação
	tx := config.DB.Begin()

	// Atualizar os registros no banco de dados
	for i, pdt := range pdts {
		//soma = somaPorsalariobaseUnidadeFuncaoPdt[pdt.Saldo_salario][pdt.Unidade][pdt.Funcao]
		pdts[i].Salariobasefo = somaPorsalariobaseUnidadeFuncaoPdt[pdt.Saldo_salario][pdt.Unidade][pdt.Funcao]

		// Atualizar os registros na transação
		if err := tx.Model(&pdts[i]).Update("Salariobasefo", pdts[i].Salariobasefo).Error; err != nil {
			// Lidar com erros, rolar de volta a transação se necessário
			tx.Rollback()
			return
		}
	}

	// Commit da transação
	tx.Commit()

	//soma = 0.0
}