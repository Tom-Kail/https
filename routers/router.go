package routers

import (
	"github.com/Tom-Kail/https/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
