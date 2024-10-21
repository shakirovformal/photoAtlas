package pkg

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4"
)

type Database struct {
	cnt int
	// dbpg_addr  string
	unique_id  string
	name_file  string
	file_exist bool
}

const dbpg_addr = "postgres://postgres:toor@localhost:5432/photoatlas"

func DB_Select(file_name string) int {
	//sql1 := "SELECT unique_id, name_file FROM images"

	sql2 := "SELECT COUNT(*) AS CNT FROM (SELECT * FROM images WHERE name_file=$1 LIMIT 1)"
	//sql3 := "SELECT row_number() OVER (ORDER BY unique_id) AS i FROM images_client_not_look t"

	ctx := context.Background()
	db, err := pgx.Connect(ctx, dbpg_addr)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	rows, err := db.Query(ctx, sql2, file_name)
	if err != nil {
		log.Fatal(err.Error())
	}
	var dbm Database
	//for j := 0; j < 1; j++ {
	for rows.Next() {

		if err := rows.Scan(&dbm.cnt); err != nil {
			log.Fatal(err)
		}

	}

	defer rows.Close()

	db.Close(ctx)
	return dbm.cnt
}

func DB_Delete() string {

	sql := "DELETE FROM images_client_not_look"

	ctx := context.Background()
	db, err := pgx.Connect(ctx, dbpg_addr)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	defer db.Close(ctx)

	rows, err := db.Query(ctx, sql)
	if err != nil {
		log.Fatalf("QueryRow failed: %v", err)
	}
	_ = rows
	return "SQL Deleted ok"
}

func DB_Insert(first_var string) string {
	sql := "INSERT INTO images (name_file) VALUES ($1)"

	ctx := context.Background()
	db, err := pgx.Connect(ctx, dbpg_addr)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	defer db.Close(ctx)

	rows, err := db.Query(ctx, sql, first_var)
	if err != nil {
		log.Fatalf("QueryRow failed: %v", err)
	}
	_ = rows
	return fmt.Sprint("SQL insertion ok. Добавили в список %v", first_var)
}

func Database_conn() {

	ctx := context.Background()
	db, err := pgx.Connect(ctx, dbpg_addr)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	defer db.Close(ctx)
	log.Println("Connect database succesful")

	//fmt.Println("Отработал метод insert у интерфейса DB")
}

func DB_UpdateNameFile(first_var string) string {
	sql := "UPDATE images SET name_file = $1 WHERE name_file = $2;"

	ctx := context.Background()
	db, err := pgx.Connect(ctx, dbpg_addr)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	defer db.Close(ctx)

	changed_name := first_var + "_unreachable"

	rows, err := db.Query(ctx, sql, changed_name, first_var)
	if err != nil {
		log.Fatalf("QueryRow failed: %v", err)
	}
	_ = rows
	return fmt.Sprint("SQL insertion ok. Добавили в список %v", first_var)
}
