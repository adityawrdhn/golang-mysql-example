package main

import (
	"apigetspot/controllers"
	_ "apigetspot/routers"
	"apigetspot/utils"

	"github.com/astaxie/beego"
)

func main() {
	utils.InitSql()
	utils.InitTemplate()
	utils.InitCache()
	utils.InitBootStrap()
	beego.ErrorController(&controllers.ErrorController{})

	beego.Run()

	// if beego.BConfig.RunMode == "dev" {
	// 	beego.BConfig.WebConfig.DirectoryIndex = true
	// 	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	// }
	// beego.Run()
}
