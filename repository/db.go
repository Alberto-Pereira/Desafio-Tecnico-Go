package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = 123
	dbname   = "desafiotecnico"
)

func StartDB() *sql.DB {
	// psqlInfo := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", user, password, user, port, dbname) // DOCKER

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%d dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Erro ao abrir conexão com banco de dados", err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Erro ao verificar se a conexão com o banco de dados ainda está ativa", err)
	}

	fmt.Println("Banco de dados conectado!")

	return db
}
