package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "meu_usuario:minha_senha@tcp(db:3306)/meu_banco")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Realizar uma consulta
	rows, err := db.Query("SELECT * FROM usuarios")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	// Iterar pelos resultados e imprimir no terminal
	for rows.Next() {
		var id int
		var usuario, senha string
		if err := rows.Scan(&id, &usuario, &senha); err != nil {
			panic(err.Error())
		}
		fmt.Printf("ID: %d, Usuário: %s, Senha: %s\n", id, usuario, senha)
	}

	// Verificar se houve algum erro durante a iteração
	if err := rows.Err(); err != nil {
		panic(err.Error())
	}
}
