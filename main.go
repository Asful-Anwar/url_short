package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main(){
dsn := "root:@tcp(127.0.0.1:3306)/url_shortener"
db, err := sql.Open("mysql", dsn)
if err !=nil {
	panic(err)
}
defer db.Close()

// Coba Koneksi
err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to the database successfully!")

}