package pkg

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4"
)

const dbpg_addr = "postgres://postgres:toor@localhost:5432/photoatlas"

func DB_Select(file_name string) bool {
	sql := "SELECT EXISTS(SELECT * FROM images_client_not_look WHERE name_file=$1)"

	ctx := context.Background()
	db, err := pgx.Connect(ctx, dbpg_addr)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	defer db.Close(ctx)
	file_name = fmt.Sprint("'" + file_name + "'")
	var temp_prov bool
	err = db.QueryRow(ctx, sql, file_name).Scan(&temp_prov)
	fmt.Printf("Проверяем наличие файла %s. Результат: %v. ", file_name, temp_prov)
	if temp_prov == true {
		fmt.Printf("Скачивать не будем\n")
	} else if temp_prov == false {
		fmt.Printf("Скачаем\n")
	}
	if err != nil {
		log.Fatalf("QueryRow failed: %v", err)
	}

	//fmt.Println(err)
	//fmt.Println("SQL Select ok")
	return temp_prov
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

func DB_Insert(first_var int, second_var string) string {
	sql := "INSERT INTO images_client_not_look VALUES ($1,$2)"

	ctx := context.Background()
	db, err := pgx.Connect(ctx, dbpg_addr)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	defer db.Close(ctx)

	rows, err := db.Query(ctx, sql, first_var, second_var)
	if err != nil {
		log.Fatalf("QueryRow failed: %v", err)
	}
	_ = rows
	return fmt.Sprint("SQL insertion ok. Добавили в список %v", second_var)
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
