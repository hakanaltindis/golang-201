package main

import (
	"fmt"
	"net/http"
)

type Human struct {
	FirstName string
	LastName  string
	Age       int
}

func (h Human) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.FirstName = "Hakan"
	h.LastName = "Altindiş"
	h.Age = 34

	r.ParseForm()

	fmt.Println(r.Form)

	fmt.Println("path:", r.URL.Path)

	fmt.Fprintf(w, "<table><tr><td><b>Name</b></td><td><b>Surname</b></td><td><b>Age</b></td></tr><tr><td>%s</td><td>%s</td><td>%d</td></tr></table>", h.FirstName, h.LastName, h.Age)
}

func main() {
	var h Human
	err := http.ListenAndServe("localhost:9000", h)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
