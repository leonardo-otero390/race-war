package tests

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"testing"

	"github.com/leonardo-otero390/race_war/controllers"
	"github.com/leonardo-otero390/race_war/database"
	"github.com/leonardo-otero390/race_war/models"
	"github.com/leonardo-otero390/race_war/seed"
	"github.com/leonardo-otero390/race_war/tests/factories"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	refreshUserTable()
	user := factories.GenUser()
	userJson, _ := json.Marshal(user)

	req := Request{http.MethodPost, "/users", bytes.NewBuffer(userJson)}

	res := MockServer(&req, controllers.CreateUser)

	var resUser models.User
	err := json.Unmarshal([]byte(res.Body.String()), &resUser)
	if err != nil {
		log.Fatalf("Cannot convert to json: %v\n", err)
	}

	assert.Equal(t, user.Email, resUser.Email)
	assert.Equal(t, user.Nickname, resUser.Nickname)
	assert.Equal(t, http.StatusOK, res.Code)
}

func TestGetUsers(t *testing.T) {
	refreshUserTable()

	_, err := seed.SeedUsers()
	if err != nil {
		log.Fatal(err)
	}

	req := Request{http.MethodGet, "/users", nil}

	res := MockServer(&req, controllers.FindUsers)

	var users []models.User
	err = json.Unmarshal([]byte(res.Body.String()), &users)
	if err != nil {
		log.Fatalf("Cannot convert to json: %v\n", err)
	}

	assert.Equal(t, len(users), 2)
	assert.Equal(t, http.StatusOK, res.Code)
}

func refreshUserTable() {
	database.ConectaComBancoDeDados()
	err := database.DB.Migrator().DropTable(&models.User{})
	if err != nil {
		log.Panic("Error to drop USER table")
	}

	err = database.DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Panic("Error to migrate USER table")
	}
	log.Printf("Successfully refreshed table")
}
