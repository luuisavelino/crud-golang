package database

import (
	"database/sql"
)

func ConectaComBancoDeDados() *sql.DB {
	conexao := "user=root dbname=loja password=root host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)

	if err != nil {
		panic(err.Error())
	}

	return db
}
