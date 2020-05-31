package main

import (
	"toriniku/db"
	"toriniku/router"
)

func main() {

	dbConn := db.Init()
	router.Router(dbConn)
}
