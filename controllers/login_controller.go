package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leonardo-otero390/race_war/controllers/auth"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/validator.v2"
)

type Credentials struct {
	Nickname string `json:"nickname" validate:"nonzero"`
	Password string `json:"password" validate:"nonzero"`
}

func Login(c *gin.Context) {
	var credentials Credentials

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	if err := validator.Validate(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	user := FindUserByNick(credentials.Nickname)
	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User ninckname doesn't exist"})
		return
	}

	if err := user.VerifyPassword(credentials.Password); err == bcrypt.ErrMismatchedHashAndPassword {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Nickname/Password doesn't match"})
		return
	}
	token, err := auth.CreateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": "Can't generate token"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"token": token,
	})

}
