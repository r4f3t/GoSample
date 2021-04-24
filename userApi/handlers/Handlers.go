package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	. "../models"
	. "../utils"
)

func Run() {
	http.HandleFunc("/", Handler)

	http.ListenAndServe(":8080", nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	//page object
	page := Page{ID: 7, Name: "Users", Description: "Gettin User List", URI: "/users"}
	users := LoadUsers()
	interests := LoadIntersts()
	interestMappings := LoadInterstMappings()

	var newUsers []User
	for _, user := range users {
		for _, interestMapping := range interestMappings {
			if user.ID == interestMapping.UserID {
				for _, interest := range interests {
					if interest.ID == interestMapping.InterestID {
						user.Interest = append(user.Interest, interest)
					}
				}
			}
		}
		newUsers = append(newUsers, user)
	}

	viewModel := UserViewModel{Page: page, Users: newUsers}
	data, _ := json.Marshal(viewModel)
	fmt.Println(data)
	w.Write([]byte(data))

}

func LoadUsers() []User {
	bytes, _ := ReadFile("../json/users.json")

	var users []User
	json.Unmarshal([]byte(bytes), &users)
	return users
}

func LoadIntersts() []Interest {
	bytes, _ := ReadFile("../json/interests.json")

	var interests []Interest
	json.Unmarshal([]byte(bytes), &interests)
	return interests
}

func LoadInterstMappings() []InterestMapping {
	bytes, _ := ReadFile("../json/userInterestMappings.json")

	var interestMapping []InterestMapping
	json.Unmarshal([]byte(bytes), &interestMapping)
	return interestMapping
}
