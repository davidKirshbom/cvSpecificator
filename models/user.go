package models

import (
	"encoding/json"
	"os"
)

type User struct {
	FirstName          string       `json:"firstName"`
	LastName           string       `json:"lastName"`
	Phone              string       `json:"phone"`
	Email              string       `json:"email"`
	Address            string       `json:"address"`
	Title              string       `json:"title"`
	About              string       `json:"about"`
	Experience         []Experience `json:"experience"`
	OptimizeExperience []Experience `json:"optimizeExperience"`
	Education          []Education `json:"education"`
	Skills             []string    `json:"skills"`
	Languages          []string    `json:"languages"`
	Link               Link        `json:"link"`
	State              State       `json:"state"`
}

func (user *User) Reset() {
	user.FirstName = ""
	user.LastName = ""
	user.Title = ""
	user.About = ""
	user.Experience = []Experience{}
	user.Education = []Education{}
	user.Skills = []string{}
	user.Languages = []string{}
	user.Link = Link{}
	user.State = State{}
}

func (user *User) Save() error {
	file, err := json.MarshalIndent(user, "", " ")
	if err != nil {
		return err
	}
	err = os.WriteFile("user.json", file, 0644)
	if err != nil {
		return err
	}
	return nil
}
func (user *User) Load() error {
	userJson, err := os.ReadFile("user.json")
	if err != nil {
		return err
	}
	json.Unmarshal(userJson, user)
	return nil
}