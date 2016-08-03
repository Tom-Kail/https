package main

import (
	_ "github.com/Tom-Kail/https/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.BConfig.Listen.HTTPSCertFile = "conf/myserver.crt"
	beego.BConfig.Listen.HTTPSKeyFile = "conf/myserver.key"
	beego.BConfig.Listen.HTTPSPort = 10034
	beego.BConfig.Listen.EnableHTTPS = true
	beego.BConfig.Listen.EnableHTTP = false

	beego.Run()
}
