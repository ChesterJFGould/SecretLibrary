package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/submitBook", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			http.ServeFile(w, r, "submitBook.html")
		case "POST":
			id := r.FormValue("id")
			title := r.FormValue("title")
			location := r.FormValue("location")
			directions := r.FormValue("directions")

			fmt.Printf("ID: %s\nTitle: %s\nLocation: %s\nDirections: %s\n\n", id, title, location, directions)

			_, err := db.Exec(fmt.Sprintf("INSERT INTO books VALUES ( \"%s\", \"%s\", \"%s\", \"%s\" );\n", id, title, location, directions));
			if err != nil {
				panic(err)
			}
		}
	})

	http.ListenAndServe(":3141", nil)
}
