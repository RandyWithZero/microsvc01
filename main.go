package main

import (
	"github.com/astaxie/beego"
	_ "microsvc01/conf"
	_ "microsvc01/filter"
	"microsvc01/routers"
)

func main() {
	routers.Router()
	beego.Run()

}
