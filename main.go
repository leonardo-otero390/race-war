package main

import (
	"github.com/leonardo-otero390/race_war/database"
	"github.com/leonardo-otero390/race_war/routes"
	"github.com/leonardo-otero390/race_war/seed"
)

func main() {
	database.ConectaComBancoDeDados()
	seed.Load(database.DB)
	routes.HandleRequest()
}
