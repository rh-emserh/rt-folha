
package controller

import (

	"example.com/gogin/config"
	"example.com/gogin/models"

//"gorm.io/gorm"


//"encoding/json"

//"sync"

)



func CalcTotalSalarioFolha() {
    //var pdt []models.PDT
	var pdts []models.PDT
	config.DB.Find(&pdts)

	var iadvhs []models.Iadvh
	config.DB.Find(&iadvhs)

somaPortotalsalariofolhaUnidadeFuncaoPdt := make(map[float64]map[string]map[string]float64)
		



for _, pdt := range pdts {
	//soma2=0.0
	if _, ok := somaPortotalsalariofolhaUnidadeFuncaoPdt[pdt.Total_salario_folha]; !ok {
		somaPortotalsalariofolhaUnidadeFuncaoPdt[pdt.Total_salario_folha] = make(map[string]map[string]float64)
	}
	if _, ok := somaPortotalsalariofolhaUnidadeFuncaoPdt[pdt.Total_salario_folha][pdt.Unidade]; !ok {
		somaPortotalsalariofolhaUnidadeFuncaoPdt[pdt.Total_salario_folha][pdt.Unidade] = make(map[string]float64)
	}
	somaPortotalsalariofolhaUnidadeFuncaoPdt[pdt.Total_salario_folha][pdt.Unidade][pdt.Funcao] =0.0

 
	somaPortotalsalariofolhaUnidadeFuncaoPdt[pdt.Total_salario_folha][pdt.Unidade][pdt.Funcao] += pdt.Salariobasefo + pdt.Auxilio_alimentacaofo+pdt.Gratificacaofo+pdt.Insalubridade+pdt.Periculosidadefo+pdt.Adicional_noturnofo+pdt.Dsrfo

	
}

for _, iadvh := range iadvhs {
	//soma1=0.0
	if _, ok := somaPortotalsalariofolhaUnidadeFuncaoPdt[iadvh.Total_salario_folha]; !ok {
		somaPortotalsalariofolhaUnidadeFuncaoPdt[iadvh.Total_salario_folha] = make(map[string]map[string]float64)
	}
	if _, ok := somaPortotalsalariofolhaUnidadeFuncaoPdt[iadvh.Total_salario_folha][iadvh.Unidade]; !ok {
		somaPortotalsalariofolhaUnidadeFuncaoPdt[iadvh.Total_salario_folha][iadvh.Unidade] = make(map[string]float64)
	} 
	//somaPortotalsalariofolhaUnidadeFuncaoPdt[iadvh.Total_salario_folha][iadvh.Unidade][iadvh.Funcao]=0.0
	somaPortotalsalariofolhaUnidadeFuncaoPdt[iadvh.Total_salario_folha][iadvh.Unidade][iadvh.Funcao] += iadvh.Salario_base+iadvh.Gratificaoia+iadvh.Insalubridade+iadvh.Periculosidadeia+iadvh.Adicional_noturnoia+iadvh.Dsria
	
 
   // soma1= somaPortotalsalariofolhaUnidadeFuncaoPdt[iadvh.Total_salario_folha][iadvh.Unidade][iadvh.Funcao]
}
	
		
		
		

		
		for i,pdt:= range pdts {
			
			pdts[i].Total_salario_folha =somaPortotalsalariofolhaUnidadeFuncaoPdt[pdt.Total_salario_folha][pdt.Unidade][pdt.Funcao]

		}
		config.DB.Save(&pdts)

}