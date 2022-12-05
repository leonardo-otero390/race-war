package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leonardo-otero390/race-war/database"
	"github.com/leonardo-otero390/race-war/models"
	"golang.org/x/crypto/bcrypt"
)

func FindUsers(c *gin.Context) {
	var users []models.User
	database.DB.Find(&users)
	c.JSON(200, users)
}

func FindUserByEmail(email string) models.User {
	var user models.User
	database.DB.Where(&models.User{Email: email}).First(&user)
	return user
}

func FindUserByNick(nick string) models.User {
	var user models.User
	database.DB.Where(&models.User{Nickname: nick}).First(&user)
	return user
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	if err := user.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	isEmailInUse := FindUserByEmail(user.Email)
	if isEmailInUse.ID != 0 {
		c.JSON(http.StatusConflict, gin.H{
			"error": "Email already taken."})
		return
	}

	isNickInUse := FindUserByNick(user.Nickname)
	if isNickInUse.ID != 0 {
		c.JSON(http.StatusConflict, gin.H{
			"error": "Nick already taken."})
		return
	}

	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": "Can't process password"})
		return
	}

	user.Password = hashedPassword

	database.DB.Create(&user)
	c.JSON(http.StatusOK, user)
}
