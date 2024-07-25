package shortenurl

import (
	"fmt"
	"math/rand"
)

func ShortenURL(url string) string {
	// create a random string of 6 characters
	s := ""
	for i := 0; i < 6; i++ {
		s += string(rand.Intn(26) + 97) // returns a random lowercase letter
	}

	shortenedURL := fmt.Sprintf("http://localhost:8080/%s", s)
	return shortenedURL
}
