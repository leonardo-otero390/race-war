package seed

import (
	"log"

	"github.com/leonardo-otero390/race-war/database"
	"github.com/leonardo-otero390/race-war/models"
	"gorm.io/gorm"
)

var users = []models.User{
	{
		Nickname: "Steven victor",
		Email:    "steven@gmail.com",
		Password: "password",
	},
	{
		Nickname: "Martin Luther",
		Email:    "luther@gmail.com",
		Password: "password",
	},
}

func SeedUsers() ([]models.User, error) {
	for i := range users {
		err := database.DB.Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			return []models.User{}, err
		}
	}
	return users, nil
}

func Load(db *gorm.DB) {
	err := db.Debug().Migrator().DropTable(&models.User{})
	if err != nil {
		log.Panic("cannot drop table: ", err)
	}
	err = db.Debug().AutoMigrate(&models.User{})
	if err != nil {
		log.Panic("cannot migrate table: ", err)
	}

	SeedUsers()
}
