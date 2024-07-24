package shortenURL

import "strings"

func Shorten(url []string) string {
	// if len(url) == 0 {
	// 	url = []string{"please enter a url"}
	// }
	// return strings.Join(url, " ")

	return "Hello" + strings.Join(url, ", ") + "!"
  }
