package main

import (
	"github.com/gregoflash05/firstapi/db"
	"github.com/gregoflash05/firstapi/routes"
)

func main() {
	db.ConnectMONGODB()
	// r := mux.NewRouter()
	routes.Startsever()
}
