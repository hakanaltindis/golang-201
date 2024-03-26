package main

/*
 * Load html file from server using template
 */

import (
	"bytes"
	"html/template"
	"net/http"
	"os"
)

type Page struct {
	Title           string
	Author          string
	Header          string
	PageDescription string
	Content         string
	URI             string
}

func loadFile(filename string) (string, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	var builder bytes.Buffer
	builder.WriteString("Go Lang 201 Educations\n")
	builder.WriteString("Topics: \n")
	builder.WriteString("1- Database Operations\n")
	builder.WriteString("2- File Operations\n")
	builder.WriteString("3- Json Operations\n")
	builder.WriteString("4- Web API Applications\n")

	uri := "https://www.github.com/hakanaltindis"
	page := Page{
		Title:           "Zero to Hero in Go",
		Author:          "Hakan Altındiş",
		Header:          "Advanced Level in Go",
		PageDescription: "Description of advanced level of Go",
		Content:         builder.String(),
		URI:             uri,
	}

	t, _ := template.ParseFiles("page.html")
	t.Execute(w, page)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9000", nil)
}
