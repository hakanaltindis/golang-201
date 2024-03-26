package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type API struct {
	Message string `json:"message"`
}

func Hello(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	id := urlParams["id"]

	messageConcat := "User ID: " + id

	message := API{Message: messageConcat}
	output, err := json.Marshal(message)
	if err != nil {
		fmt.Println("Hata oluştu")
	}
	fmt.Fprint(w, string(output))

}

func main() {
	gorillaRoute := mux.NewRouter()
	gorillaRoute.HandleFunc("/api/users/{id:[0-9]+}", Hello)
	http.Handle("/", gorillaRoute)

	http.ListenAndServe(":9000", nil)
}
