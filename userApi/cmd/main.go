package main

import (
	"errors"
	"io/ioutil"
	"net/http"

	"encoding/json"
	"fmt"
)

func main() {
	Run()
}

func Run() {
	http.HandleFunc("/getusers", Handler)

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

type Interest struct {
	ID   int
	Name string
}

type InterestMapping struct {
	UserID     int
	InterestID int
}

type User struct {
	FirstName string
	ID        int
	Interest  []Interest
	LastName  string
	Profile   string
	Username  string
}

type UserViewModel struct {
	Page  Page
	Users []User
}
type Page struct {
	ID          int
	Name        string
	Description string
	URI         string
}

func ReadFile(fileName string) (string, error) {

	if IsEmpty(fileName) {
		return "", errors.New("FileName cannot be null.")
	}

	bytes, err := ioutil.ReadFile(fileName)
	CheckErrors(err)
	return string(bytes), err
}

func CheckErrors(err error) {
	if err != nil {
		err.Error()
	}
}

func IsEmpty(value string) bool {
	return len(value) == 0
}
