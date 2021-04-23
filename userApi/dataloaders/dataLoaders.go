package dataloaders

import (
	"encoding/json"

	. "../models"
	. "../utils"
)

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
