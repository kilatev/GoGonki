package main

import (
	_ "GoGonki/gogonki_bee/test_api_creation/docs"
	_ "GoGonki/gogonki_bee/test_api_creation/routers"

	"github.com/astaxie/beego"
)

func main() {
	if beego.RunMode == "dev" {
		beego.DirectoryIndex = true
		beego.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
