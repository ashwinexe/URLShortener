package shortenurl

import (
	"net/http"

	"gorm.io/gorm"
)

type URL struct {
	ID        uint   `gorm:"primary_key"`
	Original  string `gorm:"not null"`
	Shortened string `gorm:"not null"`
}

// we'll create a new shortened URL or retrieve an existing one if it exists in DB
func CreateOrRetrieveShortenedURL(db *gorm.DB, original string) string {
	var url URL
	if err := db.First(&url, "original = ?", original).Error; err == nil {
		// if URL already exists, return the existing shortened URL
		return url.Shortened
	} 

	// if URL doesn't exist, create a new shortened URL
	shortened := ShortenURL(original)
	db.Create(&URL{Original: original, Shortened: shortened})

	return shortened}

func RedirectURL(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	// get the shortened URL from the request
	// query the database for the original URL
	// redirect the user to the original URL
	id := r.URL.Path[1:]
	var url URL
	shortened := "http://localhost:8080/" + id
	db.First(&url, "shortened = ?", shortened)
	http.Redirect(w, r, url.Original, http.StatusFound)
}
