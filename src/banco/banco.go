package banco

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB      *gorm.DB
	err     error
	CONEXAO string
)

func ConectarComBancoDeDados() {
	//CONEXAO := os.Getenv("DATABASE_URI")
	CONEXAO := "host=localhost user=root password=root dbname=feira port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(CONEXAO))
	TratarErro(err)

	err := Migrador(DB)
	TratarErro(err)
}

func TratarErro(err error) {
	if err != nil {
		log.Panic("Erro na conexão ao banco de dados: " + err.Error())
	}
}
