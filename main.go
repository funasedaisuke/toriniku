package main

import (
	"github.com/me/toriniku/db"
	"github.com/me/toriniku/router"
)

func main() {
	dbConn := db.Init()
	//routerパッケージのRouter()
	router.Router(dbConn)
}
