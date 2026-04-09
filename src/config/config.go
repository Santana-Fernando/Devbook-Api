package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	StringConexaoBanco = ""
	Porta              = 0
	SecretKey          []byte
)

func Carregar() {
	var erro error
	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Porta, erro = strconv.Atoi(os.Getenv("API_PORTA"))
	if erro != nil {
		log.Fatal(erro)
	}

	StringConexaoBanco = fmt.Sprintf("host=localhost port=5432 user=%s dbname=%s password=%s sslmode=disable",
		os.Getenv("DB_USUARIO"), os.Getenv("DB_DATABASE"), os.Getenv("DB_SENHA"))

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}
