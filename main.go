package main

import (
	"URLShortner/shortenurl"
	"fmt"
	"net/http"
)

func main() {
	// Connect to the database
	db, err := shortenurl.ConnectDb()
	if err != nil {
		panic("failed to connect database")
	}

	if err := db.AutoMigrate(&shortenurl.URL{}, &shortenurl.DomainCount{}); err != nil {
		panic("failed to migrate database schema")
	}

	db.AutoMigrate(&shortenurl.URL{})

	// Create a new shortened URL or retrieve an existing one if it exists in DB
	http.HandleFunc("/shorten", func(w http.ResponseWriter, r *http.Request) {
		original := r.FormValue("url")
		shortened := shortenurl.CreateOrRetrieveShortenedURL(db, original)
		// fmt.Println(shortened)
		fmt.Fprintf(w, `"shortened_URL": "%s"`, shortened)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		shortenurl.RedirectURL(db, w, r)
	})

	// Count the latest shortened URL and return the top 3 domains
	http.HandleFunc("/topdomains", func(w http.ResponseWriter, r *http.Request) {
		domains, err := shortenurl.GetTopDomains(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		for _, domain := range domains {
			fmt.Fprintf(w, "%s: %d\n", domain.Domain, domain.Count)
		}
	})

	http.ListenAndServe(":8080", nil)
}
