package main

import (
	"fmt"
	"os"
	"shortenURL"
)

func main() {
	if len(os.Args) >1{
		fmt.Println(shortenURL.Shorten(os.Args[1:]))
	} else {
		fmt.Println("Please enter a URL")
	}
}