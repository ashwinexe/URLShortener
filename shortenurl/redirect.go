package shortenurl

import (
	"fmt"
	"net/http"
	"net/url"
	"gorm.io/gorm"
)

type URL struct {
	ID        uint   `gorm:"primary_key"`
	Original  string `gorm:"not null;unique"`
	Shortened string `gorm:"not null;unique"`
}

type DomainCount struct {
	ID uint `gorm:"primary_key"`
	Domain string `gorm:"not null;unique"`
	Count int `gorm:"not null"`
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

	// Find and increment the domain count
	domain := extractDomain(original)
	var domainCount DomainCount
	if err := db.First(&domainCount, "domain = ?", domain).Error; err == nil {
		domainCount.Count++
		db.Save(&domainCount)
	} else if err == gorm.ErrRecordNotFound {
		db.Create(&DomainCount{Domain: domain, Count: 1})
	} else {
		return ""
	}

	return shortened
}

// Extract the domain name from the URL
func extractDomain(rawURL string) string {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return ""
	}
	fmt.Println(parsedURL.Host)
	return parsedURL.Host

}

// Count the latest shortened URL and return the top 3 domains
func GetTopDomains(db *gorm.DB) ([]DomainCount, error) {
	var domains []DomainCount
	if err := db.Order("count desc").Limit(3).Find(&domains).Error; err != nil {
		return nil, err
	}
	 return domains, nil
}

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
