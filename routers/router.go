package routers

import (
	"webconsole_sma/controllers"
	_ "webconsole_sma/utils"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("host/", &controllers.HostController{})
	beego.Router("host/ws", &controllers.WebSocketController{})
	// beego.Router("/host", &controllers.HostController{})
	beego.Router("/step", &controllers.StepController{})
	beego.Router("/step/edit", &controllers.StepController{}, "get:Edit")
	beego.AutoRouter(&controllers.ServiceController{})
	beego.Router("/service", &controllers.ServiceController{})
	beego.Router("/file/*", &controllers.FileController{})
	// beego.AutoRouter(&controllers.FileController{})
}
