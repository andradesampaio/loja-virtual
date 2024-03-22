package config


import (
    "database/sql"
    "fmt"
    exeption "loja-virtual/internal/exception"

    _ "github.com/lib/pq"
)

func Connect() *sql.DB {
	
	conexao := "user=root dbname=mydatabase password=root host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	exeption.Error(err)

    fmt.Println("Successfully connected!")
	return db
}