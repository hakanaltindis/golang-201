package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type API struct {
	Message string `json:"message"`
}

type User struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Age       int    `json:"age"`
}

const apiRoot string = "/api"

func main() {

	http.HandleFunc(apiRoot, func(w http.ResponseWriter, r *http.Request) {
		message := API{"API Home"}
		output, err := json.Marshal(message)
		checkError(err)

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, string(output))
	})

	http.HandleFunc(apiRoot+"/users", func(w http.ResponseWriter, r *http.Request) {
		users := []User{
			User{Id: 1, FirstName: "Hakan", LastName: "Altındiş", Age: 34},
			User{Id: 2, FirstName: "Sezen", LastName: "Altındiş", Age: 34},
			User{Id: 3, FirstName: "Atlas", LastName: "Altındiş", Age: 3},
		}

		output, err := json.Marshal(users)
		checkError(err)

		fmt.Fprintf(w, string(output))
	})

	http.HandleFunc(apiRoot+"/me", func(w http.ResponseWriter, r *http.Request) {
		user := User{1, "Hakan", "Altındiş", 34}

		output, err := json.Marshal(user)
		checkError(err)

		fmt.Fprintf(w, string(output))
	})

	http.ListenAndServe(":9000", nil)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal Error :", err.Error())
		os.Exit(1)
	}
}
