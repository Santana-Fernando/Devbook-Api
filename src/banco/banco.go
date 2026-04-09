package banco

import (
	"api/src/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Conectar abre a conexão com o banco de dados
func Conectar() (*gorm.DB, error) {

	db, err := gorm.Open(postgres.Open(config.StringConexaoBanco), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
