package main

import (
	"github.com/leonardo-otero390/race-war/database"
	"github.com/leonardo-otero390/race-war/routes"
	"github.com/leonardo-otero390/race-war/seed"
)

func main() {
	database.ConectaComBancoDeDados()
	seed.Load(database.DB)
	routes.HandleRequest()
}
