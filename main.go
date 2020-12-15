package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4"
)

const DATABASE_URL = "postgres://user:123@localhost:5432/todo"

func main() {
	conn, err := pgx.Connect(context.Background(), DATABASE_URL)
	if err != nil {
		log.Fatal("Connection failed\n", err)
	}
	defer conn.Close(context.Background())
	var greeting string
	err = conn.QueryRow(context.Background(), "select 'Hello, World!'").Scan(&greeting)
	if err != nil {
		log.Fatal("QueryRow failed\n", err)
	}
	fmt.Println(greeting)
}
