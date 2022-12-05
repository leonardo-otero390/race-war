package factories

import (
	"fmt"
	"log"

	"github.com/bxcodec/faker/v4"
	"github.com/leonardo-otero390/race-war/controllers"
	"github.com/leonardo-otero390/race-war/database"
)

type User struct {
	Email    string `faker:"email"`
	Nickname string `faker:"username"`
	Password string `faker:"password"`
}

func GenUser() User {
	var user User
	err := faker.FakeData(&user)
	if err != nil {
		fmt.Println(err)
	}
	return user
}

func (user *User) Insert() {
	userHashed := *user

	hashedPassword, err := controllers.HashPassword(user.Password)
	if err != nil {
		log.Fatal("Error to hash password", err)
	}

	userHashed.Password = hashedPassword

	database.DB.Create(&userHashed)
}
