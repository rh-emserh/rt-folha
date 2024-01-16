package controller

import (
	
	"example.com/gogin/config"
	"example.com/gogin/models"

//"gorm.io/gorm"

//"encoding/json"

//"sync"

)




func CalcPericulosidade() {
    //var pdt []models.PDT
	var pdts []models.PDT
	config.DB.Find(&pdts)

	var iadvhs []models.Iadvh
	config.DB.Find(&iadvhs)
	
	somaPorpericulosidadeUnidadeFuncaoPdt := make(map[float64]map[string]map[string]float64)

    for _, pdt := range pdts {
		
        if _, ok := somaPorpericulosidadeUnidadeFuncaoPdt[pdt.Periculosidade]; !ok {

            somaPorpericulosidadeUnidadeFuncaoPdt[pdt.Periculosidade] = make(map[string]map[string]float64)
        }
        if _, ok := somaPorpericulosidadeUnidadeFuncaoPdt[pdt.Periculosidade][pdt.Unidade]; !ok {
            somaPorpericulosidadeUnidadeFuncaoPdt[pdt.Periculosidade][pdt.Unidade] = make(map[string]float64)
        }
		somaPorpericulosidadeUnidadeFuncaoPdt[pdt.Periculosidade][pdt.Unidade][pdt.Funcao]=0.0
        somaPorpericulosidadeUnidadeFuncaoPdt[pdt.Periculosidade][pdt.Unidade][pdt.Funcao] += pdt.Periculosidade
    }
    
    for _, iadvh := range iadvhs {
		
        if _, ok := somaPorpericulosidadeUnidadeFuncaoPdt[iadvh.Periculosidade]; !ok {
            somaPorpericulosidadeUnidadeFuncaoPdt[iadvh.Periculosidade] = make(map[string]map[string]float64)
        }
        if _, ok := somaPorpericulosidadeUnidadeFuncaoPdt[iadvh.Periculosidade][iadvh.Unidade]; !ok {
            somaPorpericulosidadeUnidadeFuncaoPdt[iadvh.Periculosidade][iadvh.Unidade] = make(map[string]float64)
        }

        somaPorpericulosidadeUnidadeFuncaoPdt[iadvh.Periculosidade][iadvh.Unidade][iadvh.Funcao] += iadvh.Periculosidade
    }
    
    for i, pdt := range pdts {
        pdts[i].Periculosidadefo = somaPorpericulosidadeUnidadeFuncaoPdt[pdt.Periculosidade][pdt.Unidade][pdt.Funcao]
    }


	config.DB.Save(&pdts)

}