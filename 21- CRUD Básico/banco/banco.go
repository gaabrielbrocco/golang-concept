package banco

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" //driver de conexao com mysql
)

// Conectar abre a conexão com o banco
func Conectar() (*sql.DB, error) {
	stringConexao := "golang:gabi15102004@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		return nil, erro
	}

	return db, nil
}
