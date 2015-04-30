package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	// database import
	"database/sql"
	_ "github.com/lib/pq"
)

type Test struct {
	Message string
}

func handler(w http.ResponseWriter, r *http.Request) {

	db, err := sql.Open("postgres", "user=pqgotest dbname=pqgotest sslmode=verify-full")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query(`SELECT name FROM users WHERE favorite_fruit = $1
	OR age BETWEEN $2 AND $2 + 3`, "orange", 64)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	test := Test{
		fmt.Sprintf("Hi there, I love %s!", r.URL.Path[1:]),
	}

	b, err := json.Marshal(test)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, "%s", b)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
