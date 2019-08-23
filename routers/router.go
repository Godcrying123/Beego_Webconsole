package routers

import (
	"webconsole_sma/controllers"
	_ "webconsole_sma/utils"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.IndexController{})
	beego.Router("/step", &controllers.StepController{})
	beego.Router("host/", &controllers.HostController{})
	beego.Router("host/ws", &controllers.HostWebSocketController{})
	beego.Router("ssh/ws", &controllers.SSHWebSocketController{})
	// beego.Router("/step/edit", &controllers.StepController{}, "get:Edit")
	beego.Router("/service", &controllers.ServiceController{})
	beego.Router("/service/ws", &controllers.ServiceWebSocketController{})
	// beego.Router("/service/upload", &controllers.ServiceController{}, "get:Upload")
	beego.Router("/file/*", &controllers.FileController{})
	// beego.AutoRouter(&controllers.FileController{})
}
