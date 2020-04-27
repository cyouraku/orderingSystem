package main

import (
	"github.com/Biubiubiuuuu/orderingSystem/db/mysql"
	"github.com/Biubiubiuuuu/orderingSystem/helper/configHelper"
	"github.com/Biubiubiuuuu/orderingSystem/router"
)

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	mysql.InitModel()
	router := router.Init()
	router.Run(configHelper.HTTPPort)
}