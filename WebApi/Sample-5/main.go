package main

/*
 * Load html file from server
 */

import (
	"fmt"
	"net/http"
	"os"
)

func loadFile(filename string) (string, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	var body, _ = loadFile("page.html")
	fmt.Fprint(w, body)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9000", nil)
}
