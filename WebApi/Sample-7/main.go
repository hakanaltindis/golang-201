package main

/*
 * Sample Web Api Application
 */

import (
	"encoding/json"
	"net/http"
	"os"
	"text/template"

	model "github.com/hakanaltindis/golang201webapi/models"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	page := model.Page{ID: 3, Name: "Users", Description: "List of Users", URI: "/users"}
	users := loadUsers()
	interests := loadInterests()
	interestMappings := loadInterestMappings()

	var newUsers []model.User

	for _, user := range users {
		for _, mapping := range interestMappings {
			if user.Id == mapping.UserId {
				for _, interest := range interests {
					if mapping.InterestId == interest.Id {
						user.Interests = append(user.Interests, interest)
					}
				}
			}
		}
		newUsers = append(newUsers, user)
	}

	viewModel := model.UserViewModel{Page: page, Users: newUsers}

	t, _ := template.ParseFiles("template/page.html")
	t.Execute(w, viewModel)
}

func loadFileAsBytes(filename string) ([]byte, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func loadFile(filename string) (string, error) {
	bytes, err := loadFileAsBytes(filename)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func loadUsers() []model.User {
	bytes, _ := loadFileAsBytes("json/users.json")
	var users []model.User
	json.Unmarshal(bytes, &users)
	return users
}

func loadInterests() []model.Interest {
	bytes, _ := loadFileAsBytes("json/interests.json")
	var interests []model.Interest
	json.Unmarshal(bytes, &interests)
	return interests
}

func loadInterestMappings() []model.InterestMapping {
	bytes, _ := loadFileAsBytes("json/userInterestMappings.json")
	var mappings []model.InterestMapping
	json.Unmarshal(bytes, &mappings)
	return mappings
}
