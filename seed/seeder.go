package seed

import (
	"log"

	"github.com/leonardo-otero390/race_war/models"
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

func Load(db *gorm.DB) {
	err := db.Debug().Migrator().DropTable(&models.User{})
	if err != nil {
		log.Panic("cannot drop table: ", err)
	}
	err = db.Debug().AutoMigrate(&models.User{})
	if err != nil {
		log.Panic("cannot migrate table: ", err)
	}

	for i := range users {
		db.Debug().Create(&users[i])

	}
}
