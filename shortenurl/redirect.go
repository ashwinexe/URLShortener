package shortenurl

import (
	"net/http"

	"gorm.io/gorm"
)

type URL struct {
	ID           uint   `gorm:"primary_key"`
	OriginalURL  string `gorm:"not null"`
	ShortenedURL string `gorm:"not null"`
}

func RedirectURL(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	// get the shortened URL from the request
	// query the database for the original URL
	// redirect the user to the original URL
	id := r.URL.Path[1:]
	var url URL
	shortened := "http://localhost:8080/" + id
	db.First(&url, "shortened = ?", shortened)
	http.Redirect(w, r, url.OriginalURL, http.StatusFound)
}
