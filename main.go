package main

import (
	"github.com/TranHungKT/settle_up/database"
	"github.com/gin-gonic/gin"
)

func main() {
	var server = initGin()

	database.InitDB()

	server.Run()
}

func initGin() *gin.Engine {
	var r = gin.Default()

	return r
}
