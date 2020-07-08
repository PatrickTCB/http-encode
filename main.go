package main

import (
	"fmt"
	"net/url"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		rawValue := os.Args[1]
		encodedValue := url.QueryEscape(rawValue)
		fmt.Println(encodedValue)
	} else {
		fmt.Println("This program requires an argument, namely some text to decode.")
	}
}
