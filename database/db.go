package database

import (
	"log"
	"os"
	"github.com/guilhermeonrails/api-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	stringDeConexao := "host=" + os.Getenv("DATABASE_URL") + " user=" + os.Getenv("DATABASE_USER") + " password=" + os.Getenv("DATABASE_PWD") + " dbname=" + os.Getenv("DATABASE_DBNAME") + " port=" + os.Getenv("DATABASE_DBPORT")
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados: " +  "host=" + os.Getenv("DATABASE_URL") + " user=" + os.Getenv("DATABASE_USER") + " password=" + os.Getenv("DATABASE_PWD") + " dbname=" + os.Getenv("DATABASE_DBNAME") + " port=" + os.Getenv("DATABASE_DBPORT") + " sslmode=disable")
		log.Panic(nil)
	}

	DB.AutoMigrate(&models.Aluno{})
}


// host=root.c5kq0ywq2g7o.us-east-1.rds.amazonaws.com user=postgres password=123456789 dbname=root port=5432 sslmode=disable