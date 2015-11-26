package main

import (
	_ "github.com/lesscome33/dp_server/docs"
	_ "github.com/lesscome33/dp_server/routers"

	"github.com/astaxie/beego"
)

func main() {
	if beego.RunMode == "dev" {
		beego.DirectoryIndex = true
		beego.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
