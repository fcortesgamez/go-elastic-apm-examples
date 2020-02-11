package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go.elastic.co/apm/module/apmsql"
	_ "go.elastic.co/apm/module/apmsql/mysql"
	"log"
)

var db *sql.DB

type Product struct {
	ID   int64
	name string
	desc string
}

func main() {
	setupDB()
	defer func() {
		err := db.Close()
		if err != nil {
			panic(fmt.Sprintf("failed to close database. error = %s", err))
		}
	}()
	insertData()
	queryData()
}

func setupDB() {
	var err error
	db, err = apmsql.Open("mysql", "user:password@/test")
	if err != nil {
		panic(fmt.Sprintf("failed to open database. error = %s", err))
	}

	err = db.Ping()
	if err != nil {
		panic(fmt.Sprintf("failed to ping database. error = %s", err))
	}
}

func insertData() {
	stmt, err := db.Prepare("INSERT INTO product VALUES( ?, ?, ?)")
	if err != nil {
		panic(fmt.Sprintf("failed to prepare INSERT statement. error = %s", err))
	}
	defer func() {
		err = stmt.Close()
		if err != nil {
			panic(fmt.Sprintf("failed to close INSERT statement. error = %s", err))
		}
	}()

	// Insert some products
	for i := 0; i < 5; i++ {
		p := Product{
			ID:   0,
			name: fmt.Sprintf("Product %d", i),
			desc: fmt.Sprintf("Description %d", i),
		}
		_, err := stmt.Exec(p.ID, p.name, p.desc)
		if err != nil {
			panic(fmt.Sprintf("failed to insert product %+v. error = %s", p, err))
		}
	}
}

func queryData() {
	stmt, err := db.Prepare("SELECT * FROM product")
	if err != nil {
		panic(fmt.Sprintf("failed to prepare SELECT statement. error = %s", err))
	}
	defer func() {
		err = stmt.Close()
		if err != nil {
			panic(fmt.Sprintf("failed to close SELECT statement. error = %s", err))
		}
	}()

	// Query all existing products
	rows, err := stmt.Query()
	if err != nil {
		panic(fmt.Sprintf("failed to query products. error = %s", err))
	}
	for rows.Next() {
		p := Product{}
		err := rows.Scan(&p.ID, &p.name, &p.desc)
		if err != nil {
			panic(fmt.Sprintf("failed to scan product. error = %s", err))
		}
		log.Printf("Product: %+v", p)
	}
}
