package main

import (
	"URLShortner/shortenurl"
	"fmt"
	"net/http"
)

func main() {
	db, err := shortenurl.ConnectDb()
	if err != nil {
		panic("failed to connect database")
	}

	if err := db.AutoMigrate(&shortenurl.URL{}); err != nil {
		panic("failed to migrate database schema")
	}

	db.AutoMigrate(&shortenurl.URL{})

	http.HandleFunc("/shorten", func(w http.ResponseWriter, r *http.Request) {
		original := r.FormValue("url")
		shortened := shortenurl.CreateOrRetrieveShortenedURL(db, original)
		fmt.Println(shortened)

		fmt.Fprintf(w, `"shortened_URL": "%s"`, shortened)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		shortenurl.RedirectURL(db, w, r)
	})
	http.ListenAndServe(":8080", nil)
}
