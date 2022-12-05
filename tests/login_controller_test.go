package tests

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"testing"

	"github.com/leonardo-otero390/race_war/controllers"
	"github.com/leonardo-otero390/race_war/tests/factories"
	"github.com/leonardo-otero390/race_war/tests/utils"
	"github.com/stretchr/testify/assert"
)

type AuthResponse struct {
	Token string `json:"token"`
	Error string `json:"error"`
}

func TestLoginInvalidUser(t *testing.T) {
	utils.RefreshUserTable()
	user := factories.GenUser()

	credentials := controllers.Credentials{Nickname: user.Nickname, Password: user.Password}

	cJson, _ := json.Marshal(credentials)
	req := Request{http.MethodPost, "/login", bytes.NewBuffer(cJson)}

	res := MockServer(&req, controllers.Login)

	var authRes AuthResponse

	err := json.Unmarshal([]byte(res.Body.Bytes()), &authRes)
	if err != nil {
		log.Fatalf("Cannot convert to json: %v\n", err)
	}

	assert.Equal(t, authRes.Error, "User ninckname doesn't exist")
	assert.Equal(t, http.StatusNotFound, res.Code)
}

func TestLoginWrongPassword(t *testing.T) {
	utils.RefreshUserTable()
	user := factories.GenUser()
	user.Insert()

	credentials := controllers.Credentials{Nickname: user.Nickname, Password: "wrong password"}

	cJson, _ := json.Marshal(credentials)
	req := Request{http.MethodPost, "/login", bytes.NewBuffer(cJson)}

	res := MockServer(&req, controllers.Login)

	var authRes AuthResponse

	err := json.Unmarshal([]byte(res.Body.Bytes()), &authRes)
	if err != nil {
		log.Fatalf("Cannot convert to json: %v\n", err)
	}

	assert.Equal(t, authRes.Error, "Nickname/Password doesn't match")
	assert.Equal(t, http.StatusUnauthorized, res.Code)
}

func TestLoginSuccessfully(t *testing.T) {
	utils.RefreshUserTable()
	user := factories.GenUser()
	user.Insert()

	credentials := controllers.Credentials{Nickname: user.Nickname, Password: user.Password}

	cJson, _ := json.Marshal(credentials)
	req := Request{http.MethodPost, "/login", bytes.NewBuffer(cJson)}

	res := MockServer(&req, controllers.Login)

	var authRes AuthResponse

	err := json.Unmarshal([]byte(res.Body.Bytes()), &authRes)
	if err != nil {
		log.Fatalf("Cannot convert to json: %v\n", err)
	}

	assert.NotEmpty(t, authRes.Token)
	assert.Equal(t, http.StatusCreated, res.Code)
}
