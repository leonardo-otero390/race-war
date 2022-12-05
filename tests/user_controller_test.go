package tests

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"testing"

	"github.com/leonardo-otero390/race-war/controllers"
	"github.com/leonardo-otero390/race-war/models"
	"github.com/leonardo-otero390/race-war/seed"
	"github.com/leonardo-otero390/race-war/tests/factories"
	"github.com/leonardo-otero390/race-war/tests/utils"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	utils.RefreshUserTable()
	user := factories.GenUser()
	userJson, _ := json.Marshal(user)

	req := Request{http.MethodPost, "/users", bytes.NewBuffer(userJson)}

	res := MockServer(&req, controllers.CreateUser)

	var resUser models.User
	err := json.Unmarshal([]byte(res.Body.Bytes()), &resUser)
	if err != nil {
		log.Fatalf("Cannot convert to json: %v\n", err)
	}

	assert.Equal(t, user.Email, resUser.Email)
	assert.Equal(t, user.Nickname, resUser.Nickname)
	assert.Equal(t, http.StatusOK, res.Code)
}

func TestGetUsers(t *testing.T) {
	utils.RefreshUserTable()

	_, err := seed.SeedUsers()
	if err != nil {
		log.Fatal(err)
	}

	req := Request{http.MethodGet, "/users", nil}

	res := MockServer(&req, controllers.FindUsers)

	var users []models.User
	err = json.Unmarshal([]byte(res.Body.Bytes()), &users)
	if err != nil {
		log.Fatalf("Cannot convert to json: %v\n", err)
	}

	assert.Equal(t, len(users), 2)
	assert.Equal(t, http.StatusOK, res.Code)
}
