package main

import (
	"ginchat/router"
	"ginchat/util"
)

func main() {
	util.InitConfig()
	util.InitMysql()
	r := router.Router()
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
