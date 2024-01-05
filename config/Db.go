package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"example.com/gogin/models" 
)

var DB *gorm.DB

func Connect() {
	//dsn := "postgresql://joaosilva:SrPmshDxvVN6_CGX8T91Rw@wiser-snapper-1341.g8x.cockroachlabs.cloud:26257/Pdt?sslmode=verify-full"
	dsn:= "user=postgres password=evilasio2239 host=db.sgzbzoqclinxwkajkarv.supabase.co port=5432 dbname=postgres"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.PDT{})
	db.AutoMigrate(&models.Iadvh{})

	DB = db
}