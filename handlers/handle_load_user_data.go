package handlers

import (
	"encoding/json"
	"os"

	"github.com/davidKirshbom/cvSpecificator/models"
)

func HandleLoadUserData(user *models.User) (models.User, error) {
	userJson, err := os.ReadFile("user.json")
	if err != nil {
		return models.User{},err
	}
	json.Unmarshal(userJson, user)
	return *user,nil
	
}