package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/Asful-Anwar/url-shortener/internal/handler"
	"github.com/Asful-Anwar/url-shortener/internal/repository"
	"github.com/Asful-Anwar/url-shortener/internal/service"
)

func main() {
	// Koneksi Database
	dsn := "root:@tcp(127.0.0.1:3306)/url_shortener"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Connected to the database successfully!")

	// Init Layer
	repo := repository.NewLinkRepository(db)
	svc := service.NewLinkService(repo)
	h := handler.NewLinkHandler(svc)

	// Routing
	http.HandleFunc("/create", h.CreateShortLink)

	// Run server
	fmt.Println("Server is running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
