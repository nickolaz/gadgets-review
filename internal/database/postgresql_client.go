package database

import (
	"database/sql"
	"gadgets-review/internal/logs"

	_ "github.com/lib/pq"
)

type PostgreSqlClient struct {
	*sql.DB
}

func NewSqlClient(source string) *PostgreSqlClient {
	db, err := sql.Open("postgresql", source)
	if err != nil {
		logs.Log().Errorf("Cannot create db tentat: %s", err.Error())
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		logs.Log().Errorf("cannot connect to postgresql!")
	}
	return &PostgreSqlClient{db}
}
