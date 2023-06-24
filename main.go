package main

import (
	"github.com/TranHungKT/settle_up/database"
	"github.com/TranHungKT/settle_up/router"
)

func main() {

	database.InitDB()
	router.InitGin()
}
