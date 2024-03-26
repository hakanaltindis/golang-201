package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Define types
type Name struct {
	Family   string
	Personal string
}

type Email struct {
	ID      int
	Kind    string
	Address string
}

type Interest struct {
	ID   int
	Name string
}

type Person struct {
	ID        int
	FirstName string
	LastName  string
	UserName  string
	Gender    string
	Name      Name
	Email     []Email
	Interest  []Interest
}

func GetPerson(p *Person) string {
	return p.FirstName + " " + p.LastName
}

func GetPersonEmailAddress(p *Person, index int) string {
	return p.Email[index].Address
}

func GetPersonEmail(p *Person, index int) Email {
	return p.Email[index]
}

func WriteMessage(msg string) {
	fmt.Println(msg)
}

func WriteStarLine() {
	fmt.Println("****************************")
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal Error :", err.Error())
		os.Exit(1)
	}
}

func SaveJson(filename string, key interface{}) {
	outFile, err := os.Create(filename)
	checkError(err)
	defer outFile.Close()

	encoder := json.NewEncoder(outFile)
	err = encoder.Encode(key)
	checkError(err)
}

func main() {
	person := Person{
		ID:        9,
		FirstName: "Hakan",
		LastName:  "Altındiş",
		UserName:  "hakanaltindis",
		Gender:    "Male",
		Name:      Name{Family: "CS", Personal: "Hakan"},
		Email: []Email{
			{ID: 1, Kind: "Personal", Address: "hakanaltindis@hotmail.com"},
			{ID: 1, Kind: "Work", Address: "hakan.altindis@d-teknoloji.com.tr"},
		},
		Interest: []Interest{
			{ID: 1, Name: "C#"},
			{ID: 2, Name: "Azure DevOps"},
			{ID: 3, Name: "Docker"},
		},
	}

	WriteMessage("Reading Operation Started")

	WriteMessage("Personal FullName")
	WriteStarLine()
	res := GetPerson(&person)
	WriteMessage(res)
	WriteStarLine()
	WriteMessage("\n")

	WriteMessage("Personal Email With Index")
	WriteStarLine()
	resEmail := GetPersonEmailAddress(&person, 1)
	WriteMessage(resEmail)

	WriteMessage("Writing Operation Started")
	SaveJson("person.json", person)
	WriteMessage("Writing Operation Ended")
}
