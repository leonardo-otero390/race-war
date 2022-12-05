package utils

import (
	"log"

	"github.com/leonardo-otero390/race-war/database"
	"github.com/leonardo-otero390/race-war/models"
)

func RefreshUserTable() {
	database.ConectaComBancoDeDados()
	err := database.DB.Migrator().DropTable(&models.User{})
	if err != nil {
		log.Panic("Error to drop USER table")
	}

	err = database.DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Panic("Error to migrate USER table")
	}
}
